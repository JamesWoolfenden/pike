package pike

import (
	"bytes"
	_ "embed" // required for embed
	"fmt"
	"strings"
	"text/template"
)

const (
	DefaultPolicyName = "terraform_pike"
	DefaultProject    = "pike"
	DefaultRoleID     = "terraform_pike"
)

//go:embed terraform.gcppolicy.template
var policyGCPTemplate []byte

// GCPPolicy create an IAM policy.
func GCPPolicy(permissions []string) (policy string, err error) {
	test := strings.Join(permissions, "\",\n    \"")

	// GCPPolicyDetails contains the configuration for generating a GCP IAM policy
	type GCPPolicyDetails struct {
		Name        string // Custom name for the policy
		Project     string // GCP project identifier
		RoleID      string // Unique role identifier
		Permissions string // Comma-separated list of permissions
	}

	PolicyName := DefaultPolicyName
	theDetails := GCPPolicyDetails{
		Name:        PolicyName,
		Project:     DefaultProject,
		RoleID:      DefaultRoleID,
		Permissions: test,
	}
	var output bytes.Buffer

	tmpl, err := template.New("test").Parse(string(policyGCPTemplate))
	if err != nil {
		return "", fmt.Errorf("failed to parse template %w", err)
	}

	err = tmpl.Execute(&output, theDetails)

	if err != nil {
		return "", err
	}

	return output.String(), nil
}
