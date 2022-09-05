package pike

import (
	"bytes"
	_ "embed" // required for embed
	"strings"
	"text/template"
)

//go:embed terraform.gcppolicy.template
var policyGCPTemplate []byte

// GCPPolicy create an IAM policy
func GCPPolicy(permissions []string) (string, error) {
	test := strings.Join(permissions, "\",\n    \"")

	type GCPPolicyDetails struct {
		Name        string
		Project     string
		RoleID      string
		Permissions string
	}

	PolicyName := "terraform" + randSeq(8)
	theDetails := GCPPolicyDetails{PolicyName, "pike", "terraform_pike", test}

	var output bytes.Buffer
	tmpl, err := template.New("test").Parse(string(policyGCPTemplate))
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(&output, theDetails)

	if err != nil {
		panic(err)
	}
	return output.String(), nil
}
