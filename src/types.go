package pike

import (
	"encoding/json"
	"strings"
)

// SplitSet holds a provider's permissions partitioned into base and
// escalation-class subsets.
type SplitSet struct {
	Base       []string `json:"base"`
	Escalation []string `json:"escalation"`
}

// SplitPolicy is the output of --output split: each present provider has its
// permissions divided into base (safe) and escalation-class subsets.
type SplitPolicy struct {
	AWS   *SplitSet `json:"aws,omitempty"`
	GCP   *SplitSet `json:"gcp,omitempty"`
	AZURE *SplitSet `json:"azure,omitempty"`
}

// String renders the split policy as indented JSON.
func (s SplitPolicy) String() string {
	b, _ := json.MarshalIndent(s, "", "    ")
	return string(b) + "\n"
}

// Sorted is to help split out permission to the relevant auth.
type Sorted struct {
	AWS   []string
	GCP   []string
	AZURE []string
	// Plan-only (read-side) permissions for the planner SA role.
	PlanAWS   []string
	PlanGCP   []string
	PlanAZURE []string
	// Runtime permissions needed by service accounts
	RuntimeGCP   []RuntimePermission
	RuntimeAWS   []RuntimePermission
	RuntimeAZURE []RuntimePermission
	// Existing IAM bindings from Terraform code
	IAMBindings []IAMBinding
}

// RuntimePermission tracks which resource needs which runtime permissions.
type RuntimePermission struct {
	ResourceType   string   // e.g., "google_cloud_run_v2_service"
	ResourceName   string   // e.g., "app" (from resource "type" "name")
	ServiceAccount string   // e.g., "my-sa@project.iam" or "default"
	Permissions    []string // e.g., ["secretmanager.versions.access"]
}

// IAMBinding represents an IAM binding found in Terraform.
type IAMBinding struct {
	ResourceType string // e.g., "google_project_iam_member"
	ResourceName string // e.g., "app_secrets"
	Role         string // e.g., "roles/secretmanager.secretAccessor"
	Member       string // e.g., "serviceAccount:app@project.iam" or reference like "${google_service_account.app.email}"
}

// ResourceV2 is what resources get parsed into.
type ResourceV2 struct {
	TypeName       string
	Name           string
	ResourceName   string
	Provider       string
	Attributes     []string
	ServiceAccount string // The service account this resource uses (if any)
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
	AWS        AwsOutput
	GCP        string
	AZURE      string
	RuntimeGCP string
	// Plan-role output (read-only permissions for terraform plan).
	PlanAWS   AwsOutput
	PlanGCP   string
	PlanAZURE string
}

// AwsOutput structure.
type AwsOutput struct {
	JSONOut   string
	Terraform string
}

// AsString converts an object into string (legacy single-role output).
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

// AsTwoRoleString renders both the apply role and the plan role.
// For Terraform: two policy/role resource blocks.
// For JSON: {"apply": ..., "plan": ...} structure.
func (out OutputPolicy) AsTwoRoleString(format string) string {
	if strings.ToLower(format) == terraform {
		var s strings.Builder
		if out.AWS.Terraform != "" {
			s.WriteString("# apply role — full permissions for terraform apply\n")
			s.WriteString(out.AWS.Terraform)
			s.WriteString("\n")
		}
		if out.PlanAWS.Terraform != "" {
			s.WriteString("# plan role — read-only permissions for terraform plan\n")
			s.WriteString(out.PlanAWS.Terraform)
			s.WriteString("\n")
		}
		if out.GCP != "" {
			s.WriteString("# apply role\n")
			s.WriteString(out.GCP)
			s.WriteString("\n")
		}
		if out.PlanGCP != "" {
			s.WriteString("# plan role\n")
			s.WriteString(out.PlanGCP)
			s.WriteString("\n")
		}
		if out.AZURE != "" {
			s.WriteString("# apply role\n")
			s.WriteString(out.AZURE)
			s.WriteString("\n")
		}
		if out.PlanAZURE != "" {
			s.WriteString("# plan role\n")
			s.WriteString(out.PlanAZURE)
			s.WriteString("\n")
		}
		return s.String()
	}

	// JSON: wrap apply and plan under named keys.
	type twoRoleJSON struct {
		Apply json.RawMessage `json:"apply,omitempty"`
		Plan  json.RawMessage `json:"plan,omitempty"`
	}
	out2 := twoRoleJSON{}
	if out.AWS.JSONOut != "" {
		out2.Apply = json.RawMessage(strings.TrimSpace(out.AWS.JSONOut))
	}
	if out.PlanAWS.JSONOut != "" {
		out2.Plan = json.RawMessage(strings.TrimSpace(out.PlanAWS.JSONOut))
	}
	b, _ := json.MarshalIndent(out2, "", "    ")
	return string(b) + "\n"
}
