package pike

import "strings"

// Sorted is to help split out permission to the relevant auth.
type Sorted struct {
	AWS   []string
	GCP   []string
	AZURE []string
}

// ResourceV2 is what resources get parsed into.
type ResourceV2 struct {
	TypeName     string
	Name         string
	ResourceName string
	Provider     string
	Attributes   []string
}

// Policy represents and creates IAM policy structure.
type Policy struct {
	Version    string      `json:"Version"`
	Statements []Statement `json:"Statement"`
}

// Statement is the core of an IAM policy.
type Statement struct {
	Sid      string   `json:"Sid"`
	Effect   string   `json:"Effect"`
	Action   []string `json:"Action"`
	Resource []string `json:"Resource"`
}

// NewStatement constructor.
func NewStatement(sid string, effect string, action []string, resource []string) Statement {
	if effect != allow && effect != "Deny" {
		effect = "Deny" // Default to restrictive
	}

	return Statement{Sid: sid, Effect: effect, Action: action, Resource: resource}
}

// OutputPolicy is the main output type.
type OutputPolicy struct {
	AWS   AwsOutput
	GCP   string
	AZURE string
}

// AwsOutput structure.
type AwsOutput struct {
	JSONOut   string
	Terraform string
}

// AsString converts an object into string.
func (out OutputPolicy) AsString(format string) string {
	var Output string
	if strings.ToLower(format) == terraform {
		Output = out.AWS.Terraform + "\n"
	} else {
		Output = out.AWS.JSONOut + "\n"
	}

	if out.GCP != "" {
		Output = Output + out.GCP + "\n"
	}

	if out.AZURE != "" {
		Output = Output + out.AZURE + "\n"
	}

	return Output
}
