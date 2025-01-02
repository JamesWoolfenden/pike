package pike

import (
	"bytes"
	_ "embed" // required for embed
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
	if permissions == nil {
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

	PolicyName := DefaultPolicyName
	theDetails := gCPPolicyDetails{
		Name:        PolicyName,
		Project:     DefaultProject,
		RoleID:      DefaultRoleID,
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
