package pike

import (
	"bytes"
	_ "embed" //required for embed
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"text/template"
)

//go:embed terraform.policy.template
var policyTemplate []byte

// Policy creates iam policies
type Policy struct {
	Version   string `json:"Version"`
	Statement struct {
		Effect   string   `json:"Effect"`
		Action   []string `json:"Action"`
		Resource string   `json:"Resource"`
	} `json:"Statement"`
}

// NewPolicy constructor
func NewPolicy(Actions []string) Policy {
	something := Policy{}
	something.Version = "2012-10-17"
	something.Statement.Effect = "Allow"
	something.Statement.Action = Actions
	something.Statement.Resource = "*"
	return something
}

// GetPolicy creates new iam polices from a list of Permissions
func GetPolicy(actions Sorted, output string) (string, error) {
	var Permissions []string

	Permissions = append(Permissions, actions.AWS...)

	if Permissions == nil {
		return "", errors.New("no permissions detected")
	}

	//dedupe
	Permissions = unique(Permissions)

	Policy, err2 := AWSPolicy(Permissions, output)

	if err2 != nil {
		return "", err2
	}

	return Policy, nil
}

// AWSPolicy create an IAM policy
func AWSPolicy(Permissions []string, output string) (string, error) {
	Policy := NewPolicy(Permissions)
	b, err := json.MarshalIndent(Policy, "", "    ")
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	switch output {
	case "terraform", "Terraform":

		type PolicyDetails struct {
			Policy      string
			Name        string
			Path        string
			Description string
		}

		PolicyName := "terraform" + randSeq(8)
		theDetails := PolicyDetails{string(b), PolicyName, "/", "Add Description"}

		var output bytes.Buffer
		tmpl, err := template.New("test").Parse(string(policyTemplate))
		if err != nil {
			panic(err)
		}

		err = tmpl.Execute(&output, theDetails)

		if err != nil {
			panic(err)
		}
		return output.String(), nil
	default:
		return string(b) + "\n", nil
	}
}

func unique(s []string) []string {
	inResult := make(map[string]bool)
	var result []string
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	sort.Strings(result)
	return result
}
