package pike

import (
	"bytes"
	_ "embed" //required for embed
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"text/template"
)

//go:embed terraform.policy.template
var policyTemplate []byte

// NewAWSPolicy constructor
func NewAWSPolicy(Actions []string) Policy {
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

	v := reflect.ValueOf(actions)
	typeOfV := v.Type()
	values := make([]interface{}, v.NumField())

	var Policy string
	var err error
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
		switch typeOfV.Field(i).Name {
		case "AWS":
			if actions.AWS == nil {
				continue
			}
			//dedupe
			AWSPermissions := unique(actions.AWS)
			Policy, err = AWSPolicy(AWSPermissions, output)

			if err != nil {
				log.Print(err)
				continue
			}

		case "GCP":
			if actions.GCP == nil {
				continue
			}
			//dedupe
			GCPPermissions := unique(actions.GCP)
			Policy, err = GCPPolicy(GCPPermissions)
			if err != nil {
				log.Print(err)
				continue
			}
		}
	}
	if Policy == "" {
		return Policy, errors.New("no permissions found")
	}
	return Policy, nil
}

// AWSPolicy create an IAM policy
func AWSPolicy(Permissions []string, output string) (string, error) {
	Policy := NewAWSPolicy(Permissions)
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
