package pike

import "strings"

// Sorted is to help split out permission to the relevant auth
type Sorted struct {
	AWS   []string
	GCP   []string
	AZURE []string
}

// ResourceV2 is what resources get parsed into
type ResourceV2 struct {
	TypeName     string
	Name         string
	ResourceName string
	Provider     string
	Attributes   []string
}

// Policy creates iam policies
type Policy struct {
	Version    string      `json:"Version"`
	Statements []Statement `json:"Statement"`
}

// Statement is the core of an IAM policy
type Statement struct {
	Sid      string   `json:"Sid"`
	Effect   string   `json:"Effect"`
	Action   []string `json:"Action"`
	Resource string   `json:"Resource"`
}

// NewStatement constructor
func NewStatement(sid string, effect string, action []string, resource string) *Statement {
	return &Statement{Sid: sid, Effect: effect, Action: action, Resource: resource}
}

// OutputPolicy is the main output type
type OutputPolicy struct {
	AWS   AwsOutput
	GCP   string
	Azure string
}

// AwsOutput structure
type AwsOutput struct {
	JSONOut   string
	Terraform string
}

// AsString converts object into string
func (Out OutputPolicy) AsString(format string) string {
	var Output string
	if strings.ToLower(format) == "terraform" {
		Output = Out.AWS.Terraform + "\n"
	} else {
		Output = Out.AWS.JSONOut + "\n"
	}

	if Out.GCP != "" {
		Output = Output + Out.GCP + "\n"
	}

	if Out.Azure != "" {
		Output = Output + Out.Azure + "\n"
	}

	return Output
}
