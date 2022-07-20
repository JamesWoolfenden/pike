package pike

import (
	"encoding/json"
	"fmt"
)

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
func GetPolicy(actions []interface{}) {
	var Permissions []string

	for _, action := range actions {
		test := action.([]interface{})
		for _, item := range test {
			Permissions = append(Permissions, item.(string))
		}
	}

	//dedupe
	Permissions = unique(Permissions)

	Policy := NewPolicy(Permissions)
	b, err := json.MarshalIndent(Policy, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(string(b))
	fmt.Print("\n")
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
	return result
}
