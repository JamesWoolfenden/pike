package pike

import (
	"bytes"
	"context"
	_ "embed" // required for embed
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"

	"github.com/hashicorp/hcl/v2/hclsyntax"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
	"gopkg.in/ini.v1"
)

const (
	defaultPolicyName = "terraform_pike"
	defaultProject    = "pike"
	defaultRoleID     = "terraform_pike"
)

//go:embed terraform.gcppolicy.template
var policyGCPTemplate []byte

// GCPPolicy create an IAM policy.
func GCPPolicy(permissions []string, policyName string, dirName string) (string, error) {
	if len(permissions) == 0 {
		return "", &emptyPermissionsError{}
	}

	test := strings.Join(permissions, "\",\n    \"")

	// gCPPolicyDetails contains the configuration for generating a GCP IAM policy
	type gCPPolicyDetails struct {
		Name        string // Custom name for the policy
		Project     string // GCP project identifier
		RoleID      string // Unique role identifier
		Permissions string // Comma-separated list of permissions
	}

	var PolicyName string
	var roleID string

	if policyName != "" {
		PolicyName = policyName
		roleID = policyName
	} else {
		PolicyName = defaultPolicyName
		roleID = defaultRoleID
	}

	project, err := getCurrentProject(dirName)
	if err != nil {
		project = defaultProject
	}

	theDetails := gCPPolicyDetails{
		Name:        PolicyName,
		Project:     project,
		RoleID:      roleID,
		Permissions: test,
	}

	var output bytes.Buffer

	tmpl, err := template.New("test").Parse(string(policyGCPTemplate))
	if err != nil {
		return "", &templateParseError{err}
	}

	err = tmpl.Execute(&output, theDetails)
	if err != nil {
		return "", &templateExecuteError{err}
	}

	return output.String(), nil
}

func getCurrentProject(dirName string) (string, error) {
	// many different ways to ensure that a value for a GCP project is found
	for _, envVar := range []string{
		"GOOGLE_CLOUD_PROJECT",
		"CLOUDSDK_CORE_PROJECT",
		"GOOGLE_PROJECT",
		"GCLOUD_PROJECT",
		"GCP_PROJECT",
	} {
		if v := os.Getenv(envVar); v != "" {
			return v, nil
		}
	}

	// gcloud config is checked before ADC because a service account JSON key's
	// project_id reflects the SA's home project, not the operator's working project.
	if project := getGcloudConfigProject(); project != "" {
		return project, nil
	}

	ctx := context.Background()
	credentials, err := google.FindDefaultCredentials(ctx, compute.ComputeScope)
	if err == nil && credentials.ProjectID != "" {
		return credentials.ProjectID, nil
	}

	// Fall back to reading the project from the google provider block in the
	// scanned directory. This handles CI environments (e.g. WIF) where no
	// ambient credentials carry the project ID.
	if project := getProjectFromTerraformProvider(dirName); project != "" {
		return project, nil
	}

	return "", fmt.Errorf("no GCP project found: set GOOGLE_CLOUD_PROJECT or configure gcloud")
}

// getProjectFromTerraformProvider extracts the project value from a
// `provider "google" { project = "..." }` block in any .tf file under dirName.
func getProjectFromTerraformProvider(dirName string) string {
	if dirName == "" {
		return ""
	}

	files, err := GetTF(dirName)
	if err != nil {
		return ""
	}

	for _, file := range files {
		body, err := GetResourceBlocks(file)
		if err != nil {
			continue
		}

		for _, block := range body.Blocks {
			if block.Type != "provider" || len(block.Labels) == 0 || strings.ToLower(block.Labels[0]) != "google" {
				continue
			}

			projectAttr, ok := block.Body.Attributes["project"]
			if !ok {
				continue
			}

			if literalExpr, ok := projectAttr.Expr.(*hclsyntax.LiteralValueExpr); ok {
				return literalExpr.Val.AsString()
			}

			if templateExpr, ok := projectAttr.Expr.(*hclsyntax.TemplateExpr); ok {
				if len(templateExpr.Parts) == 1 {
					if lit, ok := templateExpr.Parts[0].(*hclsyntax.LiteralValueExpr); ok {
						return lit.Val.AsString()
					}
				}
			}
		}
	}

	return ""
}

func getGcloudConfigProject() string {
	var configPath string
	if runtime.GOOS != "windows" {
		configPath = filepath.Join(os.Getenv("HOME"), ".config", "gcloud", "configurations", "config_default")
	} else {
		configPath = filepath.Join(os.Getenv("APPDATA"), "gcloud", "configurations", "config_default")
	}

	config, err := ini.Load(configPath)
	if err != nil {
		return ""
	}

	return config.Section("core").Key("project").String()
}
