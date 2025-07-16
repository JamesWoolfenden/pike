package pike

import (
	"bytes"
	"context"
	_ "embed" // required for embed
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"

	//"github.com/go-git/go-git/v5/plumbing/format/config"
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
func GCPPolicy(permissions []string, policyName string) (string, error) {
	if permissions == nil || len(permissions) == 0 {
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

	if policyName != "" {
		PolicyName = policyName
	} else {
		PolicyName = defaultPolicyName
	}

	project, err := getCurrentProject()
	if err != nil {
		project = defaultProject
	}

	theDetails := gCPPolicyDetails{
		Name:        PolicyName,
		Project:     project,
		RoleID:      PolicyName,
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

func getCurrentProject() (string, error) {
	//many different ways to ensure that a value for a GCP project is found
	if os.Getenv("GOOGLE_CLOUD_PROJECT") != "" {
		return os.Getenv("GOOGLE_CLOUD_PROJECT"), nil
	}

	if os.Getenv("GOOGLE_PROJECT") != "" {
		return os.Getenv("GOOGLE_PROJECT"), nil
	}

	if os.Getenv("GCP_PROJECT") != "" {
		return os.Getenv("GCP_PROJECT"), nil
	}

	ctx := context.Background()
	credentials, err := google.FindDefaultCredentials(ctx, compute.ComputeScope)

	var configPath string
	if err != nil || credentials.ProjectID == "" {
		//gcloud info --format='value(config.paths.global_config_dir)'
		if runtime.GOOS != "windows" {
			configPath = filepath.Join(os.Getenv("HOME"), ".config", "gcloud", "configurations", "config_default")
		} else {
			configPath = filepath.Join(os.Getenv("APPDATA"), "gcloud", "configurations", "config_default")
		}

		config, err := ini.Load(configPath)

		if err != nil {
			return "", err
		}

		projectID := config.Section("core").Key("project").String()

		return projectID, nil
	}
	result := credentials.ProjectID

	return result, nil

}
