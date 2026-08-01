package validate_test

import (
	"testing"

	"github.com/jameswoolfenden/pike/internal/bootstrap/config"
	"github.com/jameswoolfenden/pike/internal/bootstrap/render"
	"github.com/jameswoolfenden/pike/internal/bootstrap/transform"
	"github.com/jameswoolfenden/pike/internal/bootstrap/validate"
	"github.com/jameswoolfenden/pike/internal/policy"
)

func testConfig() *config.Config {
	return &config.Config{
		AWS: config.AWSConfig{AccountID: "123456789012", Partition: "aws"},
		Principal: config.PrincipalConfig{
			Type: "iam-user",
			Name: "cli-user",
			MFA:  config.MFAConfig{Required: true, MaxAgeSeconds: 3600},
		},
		DeployRole:   config.DeployRoleConfig{Name: "TerraformDeployRole", Path: "/terraform-deploy/"},
		ManagedRoles: config.ManagedRolesConfig{Path: "/terraform-managed/"},
		Policies: config.PoliciesConfig{
			DeployBoundary: config.PolicyConfig{Path: "/boundaries/"},
			WorkloadBoundary: config.WorkloadBoundaryConfig{
				Name:           "TerraformWorkloadBoundary",
				Path:           "/boundaries/",
				Mode:           "allowlist",
				AllowedActions: []string{"logs:PutLogEvents"},
			},
		},
		PassRole:           config.PassRoleConfig{AllowedServices: []string{"ecs-tasks.amazonaws.com"}},
		ServiceLinkedRoles: config.ServiceLinkedRolesConfig{AllowedServices: []string{"ecs.amazonaws.com"}},
		Security: config.SecurityConfig{
			FailOnUnknownIAMAction:       true,
			DenyBoundaryRemoval:          true,
			DenyBoundaryPolicyMutation:   true,
			DenyHumanPrincipalManagement: true,
			DenyAccessKeyManagement:      true,
			DenyOrganizationManagement:   true,
		},
	}
}

// validArtifacts builds a full, correctly-generated set of artifacts
// using the real transform/render code, so the "everything passes"
// baseline is exercised the same way a real generate run would produce
// it, not hand-crafted to trivially satisfy the checks.
func validArtifacts(t *testing.T, cfg *config.Config) validate.Artifacts {
	t.Helper()

	p := policy.Policy{Statements: []policy.Statement{{
		Actions: []string{
			"iam:CreateRole", "iam:PutRolePolicy", "iam:PassRole",
			"iam:CreateServiceLinkedRole", "iam:GetRole",
		},
		Resources: []string{"*"},
	}}}

	boundary, err := transform.DeployBoundary(p, cfg)
	if err != nil {
		t.Fatalf("DeployBoundary() error = %v", err)
	}

	trust, err := render.TrustPolicy(cfg)
	if err != nil {
		t.Fatalf("TrustPolicy() error = %v", err)
	}

	return validate.Artifacts{
		DeployBoundary:   boundary,
		TrustPolicy:      trust,
		AssumeRolePolicy: render.AssumeRolePolicy(cfg),
	}
}

func TestValidate_ValidArtifactsHaveNoBlockingFindings(t *testing.T) {
	t.Parallel()

	cfg := testConfig()
	findings := validate.Validate(validArtifacts(t, cfg), cfg)

	for _, f := range findings {
		if f.Severity == validate.Blocking {
			t.Errorf("unexpected blocking finding on a validly generated artifact set: %+v", f)
		}
	}
}

func codes(findings []validate.Finding) []string {
	out := make([]string, len(findings))
	for i, f := range findings {
		out[i] = f.Code
	}

	return out
}

func containsCode(findings []validate.Finding, code string) bool {
	for _, f := range findings {
		if f.Code == code {
			return true
		}
	}

	return false
}

func TestValidate_RMP_E001_UnconditionalCreateRole(t *testing.T) {
	t.Parallel()

	cfg := testConfig()
	a := validArtifacts(t, cfg)

	for i, s := range a.DeployBoundary.Statements {
		if containsActionForTest(s.Action, "iam:CreateRole") {
			a.DeployBoundary.Statements[i].Condition = nil
		}
	}

	findings := validate.Validate(a, cfg)
	if !containsCode(findings, "RMP-E001") {
		t.Errorf("Validate() = %v, want RMP-E001", codes(findings))
	}
}

func TestValidate_RMP_E002_WrongBoundaryCondition(t *testing.T) {
	t.Parallel()

	cfg := testConfig()
	a := validArtifacts(t, cfg)

	for i, s := range a.DeployBoundary.Statements {
		if containsActionForTest(s.Action, "iam:CreateRole") {
			a.DeployBoundary.Statements[i].Condition = &render.Condition{
				StringEquals: map[string]any{"iam:PermissionsBoundary": "arn:aws:iam::999999999999:policy/wrong/Boundary"},
			}
		}
	}

	findings := validate.Validate(a, cfg)
	if !containsCode(findings, "RMP-E002") {
		t.Errorf("Validate() = %v, want RMP-E002", codes(findings))
	}
}

func TestValidate_RMP_E003_PassRoleWildcardResource(t *testing.T) {
	t.Parallel()

	cfg := testConfig()
	a := validArtifacts(t, cfg)

	for i, s := range a.DeployBoundary.Statements {
		if containsActionForTest(s.Action, "iam:PassRole") {
			a.DeployBoundary.Statements[i].Resource = []string{"*"}
		}
	}

	findings := validate.Validate(a, cfg)
	if !containsCode(findings, "RMP-E003") {
		t.Errorf("Validate() = %v, want RMP-E003", codes(findings))
	}
}

func TestValidate_RMP_E004_PassRoleMissingCondition(t *testing.T) {
	t.Parallel()

	cfg := testConfig()
	a := validArtifacts(t, cfg)

	for i, s := range a.DeployBoundary.Statements {
		if containsActionForTest(s.Action, "iam:PassRole") {
			a.DeployBoundary.Statements[i].Condition = nil
		}
	}

	findings := validate.Validate(a, cfg)
	if !containsCode(findings, "RMP-E004") {
		t.Errorf("Validate() = %v, want RMP-E004", codes(findings))
	}
}

func TestValidate_RMP_E005_ServiceLinkedRoleMissingCondition(t *testing.T) {
	t.Parallel()

	cfg := testConfig()
	a := validArtifacts(t, cfg)

	for i, s := range a.DeployBoundary.Statements {
		if containsActionForTest(s.Action, "iam:CreateServiceLinkedRole") {
			a.DeployBoundary.Statements[i].Condition = nil
		}
	}

	findings := validate.Validate(a, cfg)
	if !containsCode(findings, "RMP-E005") {
		t.Errorf("Validate() = %v, want RMP-E005", codes(findings))
	}
}

func TestValidate_RMP_E006_MissingBoundaryDenies(t *testing.T) {
	t.Parallel()

	cfg := testConfig()
	a := validArtifacts(t, cfg)

	var kept []render.Statement

	for _, s := range a.DeployBoundary.Statements {
		if s.Effect == "Deny" {
			continue
		}

		kept = append(kept, s)
	}

	a.DeployBoundary.Statements = kept

	findings := validate.Validate(a, cfg)
	if !containsCode(findings, "RMP-E006") {
		t.Errorf("Validate() = %v, want RMP-E006", codes(findings))
	}
}

func TestValidate_RMP_E007_UnclassifiedAction(t *testing.T) {
	t.Parallel()

	cfg := testConfig()
	a := validArtifacts(t, cfg)

	a.DeployBoundary.Statements = append(a.DeployBoundary.Statements, render.Statement{
		Sid: "Hand-edited", Effect: "Allow", Action: []string{"iam:UpdateRoleDescription"}, Resource: []string{"*"},
	})

	findings := validate.Validate(a, cfg)
	if !containsCode(findings, "RMP-E007") {
		t.Errorf("Validate() = %v, want RMP-E007", codes(findings))
	}
}

func TestValidate_RMP_E008_TrustPolicyMissingMFA(t *testing.T) {
	t.Parallel()

	cfg := testConfig()
	a := validArtifacts(t, cfg)
	a.TrustPolicy.Statement[0].Condition = nil

	findings := validate.Validate(a, cfg)
	if !containsCode(findings, "RMP-E008") {
		t.Errorf("Validate() = %v, want RMP-E008", codes(findings))
	}
}

func TestValidate_RMP_E009_AssumeRolePolicyTooBroad(t *testing.T) {
	t.Parallel()

	cfg := testConfig()
	a := validArtifacts(t, cfg)
	a.AssumeRolePolicy.Statements[0].Action = append(a.AssumeRolePolicy.Statements[0].Action, "iam:PassRole")

	findings := validate.Validate(a, cfg)
	if !containsCode(findings, "RMP-E009") {
		t.Errorf("Validate() = %v, want RMP-E009", codes(findings))
	}
}

func TestValidate_WorkloadAllowlistWildcardWarning(t *testing.T) {
	t.Parallel()

	cfg := testConfig()
	cfg.Policies.WorkloadBoundary.AllowedActions = []string{"s3:*"}

	findings := validate.Validate(validArtifacts(t, cfg), cfg)

	found := false

	for _, f := range findings {
		if f.Code == "WORKLOAD_ALLOWLIST_WILDCARD" {
			found = true

			if f.Severity != validate.Warning {
				t.Errorf("Severity = %v, want Warning", f.Severity)
			}
		}
	}

	if !found {
		t.Errorf("Validate() = %v, want WORKLOAD_ALLOWLIST_WILDCARD", codes(findings))
	}
}

func containsActionForTest(actions []string, action string) bool {
	for _, a := range actions {
		if a == action {
			return true
		}
	}

	return false
}
