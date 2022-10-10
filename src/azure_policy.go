package pike

import (
	"bytes"
	_ "embed" // required for embed
	"strings"
	"text/template"
)

//go:embed terraform.azurepolicy.template
var policyAZURETemplate []byte

// AZUREPolicy creates an Azure role definition
func AZUREPolicy(permissions []string) (string, error) {
	test := strings.Join(permissions, "\",\n    \"")

	type AzurePolicyDetails struct {
		Name        string
		Permissions string
	}

	policyName := strings.ToLower("terraform_pike" + randSeq(8))

	theDetails := AzurePolicyDetails{policyName, test}

	var output bytes.Buffer
	tmpl, err := template.New("test").Parse(string(policyAZURETemplate))
	if err != nil {
		return "", err
	}

	err = tmpl.Execute(&output, theDetails)

	if err != nil {
		return "", err
	}

	return output.String(), nil
}
