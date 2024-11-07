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
func AZUREPolicy(permissions []string) (string, error) {
	test := strings.Join(permissions, "\",\n    \"")

	type azurePolicyDetails struct {
		Name        string
		Permissions string
	}

	policyName := "terraform_pike"

	theDetails := azurePolicyDetails{policyName, test}

	var output bytes.Buffer

	tmpl, err := template.New("test").Parse(string(policyAZURETemplate))

	if err != nil {
		return "", fmt.Errorf("failed to create template %w", err)
	}

	err = tmpl.Execute(&output, theDetails)

	if err != nil {
		return "", fmt.Errorf("failed to execute template %w", err)
	}

	return output.String(), nil
}
