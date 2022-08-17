package pike

import (
	"bytes"
	_ "embed" //required for embed
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"text/template"
)

//go:embed terraform.policy.template
var policyTemplate []byte

// NewPolicy constructor
func NewPolicy(Actions []string) Policy {
	something := Policy{}
	something.Version = "2012-10-17"

	sort.Strings(Actions)

	var categories []string
	for _, action := range Actions {
		prefix := strings.Split(action, ":")[0]
		categories = append(categories, prefix)
	}

	sections := unique(categories)
	var statements []Statement

	for count, section := range sections {
		var myActions []string
		for _, action := range Actions {
			mySection := section + ":"
			if strings.Contains(action, mySection) {
				myActions = append(myActions, action)
			}
		}
		var state Statement

		state.Effect = "Allow"
		state.Sid = "VisualEditor" + strconv.Itoa(count)
		state.Action = myActions
		state.Resource = "*"
		statements = append(statements, state)
	}

	something.Statements = statements
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
