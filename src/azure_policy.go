package pike

import (
	"bytes"
	_ "embed" // required for embed
	"fmt"
	"strings"
	"text/template"
)

//go:embed terraform.azurepolicy.template
var policyAZURETemplate []byte

// AZUREPolicy creates an Azure role definition.
// permissions: slice of Azure permission strings in format "action:resource"
// Returns the policy definition as a string or an error if generation fails.
func AZUREPolicy(permissions []string, policyName string) (string, error) {
	// Add validation for empty permissions slice
	if len(permissions) == 0 {
		return "", fmt.Errorf("permissions list cannot be empty")
	}

	test := strings.Join(permissions, "\",\n    \"")

	type AzurePolicyDetails struct {
		Name        string `json:"name"`
		Permissions string `json:"permissions"`
	}

	if policyName == "" {
		policyName = DefaultPolicyName
	}

	theDetails := AzurePolicyDetails{policyName, test}

	var output bytes.Buffer

	tmpl, err := template.New("test").Parse(string(policyAZURETemplate))

	if err != nil {
		return "", &templateParseError{err}
	}

	err = tmpl.Execute(&output, theDetails)

	if err != nil {
		return "", &templateExecuteError{err}
	}

	return output.String(), nil
}
