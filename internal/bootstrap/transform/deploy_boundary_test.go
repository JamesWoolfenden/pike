package transform_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/qj0r9j0vc2/rampart/internal/bootstrap/config"
	"github.com/qj0r9j0vc2/rampart/internal/bootstrap/transform"
	"github.com/qj0r9j0vc2/rampart/internal/policy"
)

// testConfig returns the spec section 7 reference config, so expected
// ARNs in these tests match the ones spec section 12's examples show.
func testConfig() *config.Config {
	return &config.Config{
		Version: 1,
		AWS: config.AWSConfig{
			AccountID: "123456789012",
			Partition: "aws",
			Region:    "ap-northeast-2",
		},
		Principal: config.PrincipalConfig{
			MFA: config.MFAConfig{Required: true, Serial: "arn:aws:iam::123456789012:mfa/cli-user"},
		},
		DeployRole: config.DeployRoleConfig{Path: "/terraform-deploy/"},
		ManagedRoles: config.ManagedRolesConfig{
			Path: "/terraform-managed/",
		},
		Policies: config.PoliciesConfig{
			Deploy:         config.PolicyConfig{Name: "TerraformDeployPolicy", Path: "/terraform-deploy/"},
			DeployBoundary: config.PolicyConfig{Name: "TerraformDeployBoundary", Path: "/boundaries/"},
			WorkloadBoundary: config.WorkloadBoundaryConfig{
				Name: "TerraformWorkloadBoundary",
				Path: "/boundaries/",
				Mode: "allowlist",
			},
		},
		PassRole:           config.PassRoleConfig{AllowedServices: []string{"ecs-tasks.amazonaws.com", "scheduler.amazonaws.com"}},
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

func statementBySid(t *testing.T, doc []byte, sid string) map[string]any {
	t.Helper()

	var decoded struct {
		Statement []map[string]any `json:"Statement"`
	}

	if err := json.Unmarshal(doc, &decoded); err != nil {
		t.Fatalf("unmarshal document: %v", err)
	}

	for _, s := range decoded.Statement {
		if s["Sid"] == sid {
			return s
		}
	}

	t.Fatalf("no statement with Sid %q in %s", sid, doc)

	return nil
}

func TestDeployBoundary_CreateRole(t *testing.T) {
	t.Parallel()

	p := policy.Policy{Statements: []policy.Statement{{Actions: []string{"iam:CreateRole"}, Resources: []string{"*"}}}}

	doc, err := transform.DeployBoundary(p, testConfig())
	if err != nil {
		t.Fatalf("DeployBoundary() error = %v", err)
	}

	raw, err := json.Marshal(doc)
	if err != nil {
		t.Fatal(err)
	}

	s := statementBySid(t, raw, "CreateRoleWithRequiredBoundary")

	if s["Effect"] != "Allow" {
		t.Errorf("Effect = %v, want Allow", s["Effect"])
	}

	action, _ := s["Action"].([]any)
	if len(action) != 1 || action[0] != "iam:CreateRole" {
		t.Errorf("Action = %v, want [iam:CreateRole]", s["Action"])
	}

	resource, _ := s["Resource"].([]any)
	if len(resource) != 1 || resource[0] != "arn:aws:iam::123456789012:role/terraform-managed/*" {
		t.Errorf("Resource = %v, want [arn:aws:iam::123456789012:role/terraform-managed/*]", s["Resource"])
	}

	condition, _ := s["Condition"].(map[string]any)
	stringEquals, _ := condition["StringEquals"].(map[string]any)

	if stringEquals["iam:PermissionsBoundary"] != "arn:aws:iam::123456789012:policy/boundaries/TerraformWorkloadBoundary" {
		t.Errorf("Condition.StringEquals[iam:PermissionsBoundary] = %v, want the workload boundary ARN", stringEquals["iam:PermissionsBoundary"])
	}
}

func TestDeployBoundary_ManageTerraformRoles(t *testing.T) {
	t.Parallel()

	// Every action spec section 12.2 lists except PutRolePermissionsBoundary,
	// to prove that one gets added even when Pike didn't separately detect it.
	p := policy.Policy{Statements: []policy.Statement{{Actions: []string{
		"iam:AttachRolePolicy",
		"iam:DeleteRole",
		"iam:DeleteRolePolicy",
		"iam:DetachRolePolicy",
		"iam:PutRolePolicy",
		"iam:TagRole",
		"iam:UntagRole",
	}, Resources: []string{"*"}}}}

	doc, err := transform.DeployBoundary(p, testConfig())
	if err != nil {
		t.Fatalf("DeployBoundary() error = %v", err)
	}

	raw, err := json.Marshal(doc)
	if err != nil {
		t.Fatal(err)
	}

	s := statementBySid(t, raw, "ManageTerraformRoles")

	want := []string{
		"iam:AttachRolePolicy",
		"iam:DeleteRole",
		"iam:DeleteRolePolicy",
		"iam:DetachRolePolicy",
		"iam:PutRolePermissionsBoundary",
		"iam:PutRolePolicy",
		"iam:TagRole",
		"iam:UntagRole",
	}

	action, _ := s["Action"].([]any)
	if len(action) != len(want) {
		t.Fatalf("Action = %v, want %v", action, want)
	}

	for i, a := range want {
		if action[i] != a {
			t.Errorf("Action[%d] = %v, want %q", i, action[i], a)
		}
	}

	resource, _ := s["Resource"].([]any)
	if len(resource) != 1 || resource[0] != "arn:aws:iam::123456789012:role/terraform-managed/*" {
		t.Errorf("Resource = %v, want the managed-role ARN", s["Resource"])
	}
}

func TestDeployBoundary_PassRole(t *testing.T) {
	t.Parallel()

	p := policy.Policy{Statements: []policy.Statement{{Actions: []string{"iam:PassRole"}, Resources: []string{"*"}}}}

	doc, err := transform.DeployBoundary(p, testConfig())
	if err != nil {
		t.Fatalf("DeployBoundary() error = %v", err)
	}

	raw, _ := json.Marshal(doc)
	s := statementBySid(t, raw, "PassManagedRolesToApprovedServices")

	condition, _ := s["Condition"].(map[string]any)
	stringEquals, _ := condition["StringEquals"].(map[string]any)
	services, _ := stringEquals["iam:PassedToService"].([]any)

	if len(services) != 2 || services[0] != "ecs-tasks.amazonaws.com" || services[1] != "scheduler.amazonaws.com" {
		t.Errorf("Condition.StringEquals[iam:PassedToService] = %v, want the two configured services sorted", services)
	}
}

func TestDeployBoundary_PassRole_MissingAllowedServices(t *testing.T) {
	t.Parallel()

	p := policy.Policy{Statements: []policy.Statement{{Actions: []string{"iam:PassRole"}, Resources: []string{"*"}}}}

	cfg := testConfig()
	cfg.PassRole.AllowedServices = nil

	_, err := transform.DeployBoundary(p, cfg)

	var target *transform.PassRoleServicesRequiredError
	if !errors.As(err, &target) {
		t.Fatalf("DeployBoundary() error = %v, want *PassRoleServicesRequiredError", err)
	}
}

func TestDeployBoundary_ServiceLinkedRole(t *testing.T) {
	t.Parallel()

	p := policy.Policy{Statements: []policy.Statement{{Actions: []string{"iam:CreateServiceLinkedRole"}, Resources: []string{"*"}}}}

	doc, err := transform.DeployBoundary(p, testConfig())
	if err != nil {
		t.Fatalf("DeployBoundary() error = %v", err)
	}

	raw, _ := json.Marshal(doc)
	s := statementBySid(t, raw, "CreateApprovedServiceLinkedRoles")

	resource, _ := s["Resource"].([]any)
	if len(resource) != 1 || resource[0] != "*" {
		t.Errorf("Resource = %v, want [*]", s["Resource"])
	}

	condition, _ := s["Condition"].(map[string]any)
	stringEquals, _ := condition["StringEquals"].(map[string]any)
	services, _ := stringEquals["iam:AWSServiceName"].([]any)

	if len(services) != 1 || services[0] != "ecs.amazonaws.com" {
		t.Errorf("Condition.StringEquals[iam:AWSServiceName] = %v, want [ecs.amazonaws.com]", services)
	}
}

func TestDeployBoundary_ServiceLinkedRole_MissingAllowedServices(t *testing.T) {
	t.Parallel()

	p := policy.Policy{Statements: []policy.Statement{{Actions: []string{"iam:CreateServiceLinkedRole"}, Resources: []string{"*"}}}}

	cfg := testConfig()
	cfg.ServiceLinkedRoles.AllowedServices = nil

	_, err := transform.DeployBoundary(p, cfg)

	var target *transform.ServiceLinkedRoleServicesRequiredError
	if !errors.As(err, &target) {
		t.Fatalf("DeployBoundary() error = %v, want *ServiceLinkedRoleServicesRequiredError", err)
	}
}

func TestDeployBoundary_MandatoryDeniesAlwaysPresent(t *testing.T) {
	t.Parallel()

	// Empty scanned policy: the denies must still appear, driven by
	// config alone (spec section 12.5 / 4.2), not by what was detected.
	doc, err := transform.DeployBoundary(policy.Policy{}, testConfig())
	if err != nil {
		t.Fatalf("DeployBoundary() error = %v", err)
	}

	raw, _ := json.Marshal(doc)

	for _, tt := range []struct {
		sid    string
		action []string
	}{
		{"DenyWorkloadBoundaryRemoval", []string{"iam:DeleteRolePermissionsBoundary"}},
		{"DenyBoundaryPolicyMutation", []string{"iam:CreatePolicyVersion", "iam:DeletePolicy", "iam:DeletePolicyVersion", "iam:SetDefaultPolicyVersion"}},
		{"DenyOrganizationManagement", []string{"account:*", "organizations:*"}},
	} {
		s := statementBySid(t, raw, tt.sid)

		if s["Effect"] != "Deny" {
			t.Errorf("%s: Effect = %v, want Deny", tt.sid, s["Effect"])
		}

		action, _ := s["Action"].([]any)
		if len(action) != len(tt.action) {
			t.Fatalf("%s: Action = %v, want %v", tt.sid, action, tt.action)
		}

		for i, a := range tt.action {
			if action[i] != a {
				t.Errorf("%s: Action[%d] = %v, want %q", tt.sid, i, action[i], a)
			}
		}
	}

	boundaryDeny := statementBySid(t, raw, "DenyBoundaryPolicyMutation")
	resource, _ := boundaryDeny["Resource"].([]any)

	if len(resource) != 1 || resource[0] != "arn:aws:iam::123456789012:policy/boundaries/*" {
		t.Errorf("DenyBoundaryPolicyMutation Resource = %v, want [arn:aws:iam::123456789012:policy/boundaries/*]", resource)
	}
}

func TestDeployBoundary_MandatoryDeniesRespectSecurityFlags(t *testing.T) {
	t.Parallel()

	cfg := testConfig()
	cfg.Security.DenyOrganizationManagement = false

	doc, err := transform.DeployBoundary(policy.Policy{}, cfg)
	if err != nil {
		t.Fatalf("DeployBoundary() error = %v", err)
	}

	for _, s := range doc.Statements {
		if s.Sid == "DenyOrganizationManagement" {
			t.Fatal("DenyOrganizationManagement present despite security.deny_organization_management = false")
		}
	}
}

func TestDeployBoundary_UnknownIAMActionFailsClosedByDefault(t *testing.T) {
	t.Parallel()

	p := policy.Policy{Statements: []policy.Statement{{Actions: []string{"iam:UpdateRoleDescription"}, Resources: []string{"*"}}}}

	_, err := transform.DeployBoundary(p, testConfig())

	var target *transform.UnknownIAMActionError
	if !errors.As(err, &target) {
		t.Fatalf("DeployBoundary() error = %v, want *UnknownIAMActionError", err)
	}
}

func TestDeployBoundary_UnknownIAMActionAllowedWhenFlagDisabled(t *testing.T) {
	t.Parallel()

	p := policy.Policy{Statements: []policy.Statement{{Actions: []string{"iam:UpdateRoleDescription"}, Resources: []string{"*"}}}}

	cfg := testConfig()
	cfg.Security.FailOnUnknownIAMAction = false

	doc, err := transform.DeployBoundary(p, cfg)
	if err != nil {
		t.Fatalf("DeployBoundary() error = %v, want nil", err)
	}

	for _, s := range doc.Statements {
		for _, a := range s.Action {
			if a == "iam:UpdateRoleDescription" {
				t.Fatalf("unclassified action leaked into statement %q — it must be dropped, not silently allowed", s.Sid)
			}
		}
	}
}

func TestDeployBoundary_NonIAMActionsPassThrough(t *testing.T) {
	t.Parallel()

	p := policy.Policy{Statements: []policy.Statement{{Actions: []string{"ec2:RunInstances", "ec2:DescribeInstances"}, Resources: []string{"*"}}}}

	doc, err := transform.DeployBoundary(p, testConfig())
	if err != nil {
		t.Fatalf("DeployBoundary() error = %v", err)
	}

	var found []string

	for _, s := range doc.Statements {
		if s.Effect == "Allow" {
			found = append(found, s.Action...)
		}
	}

	for _, want := range []string{"ec2:RunInstances", "ec2:DescribeInstances"} {
		ok := false

		for _, a := range found {
			if a == want {
				ok = true

				break
			}
		}

		if !ok {
			t.Errorf("non-IAM action %q missing from the boundary", want)
		}
	}
}

func TestDeployBoundary_CreateRoleNeverUnconditional(t *testing.T) {
	t.Parallel()

	p := policy.Policy{Statements: []policy.Statement{{Actions: []string{"iam:CreateRole"}, Resources: []string{"*"}}}}

	doc, err := transform.DeployBoundary(p, testConfig())
	if err != nil {
		t.Fatalf("DeployBoundary() error = %v", err)
	}

	for _, s := range doc.Statements {
		for _, a := range s.Action {
			if a == "iam:CreateRole" && s.Condition == nil {
				t.Fatal("an iam:CreateRole statement has no Condition — spec section 12.1: this must never remain unconditional")
			}
		}
	}
}

func TestDeployBoundary_Deterministic(t *testing.T) {
	t.Parallel()

	p := policy.Policy{Statements: []policy.Statement{
		{Actions: []string{"iam:CreateRole", "iam:PassRole", "iam:CreateServiceLinkedRole", "iam:GetRole"}, Resources: []string{"*"}},
		{Actions: []string{"s3:GetObject", "s3:PutObject"}, Resources: []string{"*"}},
	}}

	cfg := testConfig()

	first, err := transform.DeployBoundary(p, cfg)
	if err != nil {
		t.Fatal(err)
	}

	firstJSON, err := json.Marshal(first)
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		doc, err := transform.DeployBoundary(p, cfg)
		if err != nil {
			t.Fatal(err)
		}

		got, err := json.Marshal(doc)
		if err != nil {
			t.Fatal(err)
		}

		if string(got) != string(firstJSON) {
			t.Fatalf("run %d: DeployBoundary() rendered differently:\n%s\nvs\n%s", i, got, firstJSON)
		}
	}
}

func TestDeployBoundary_NonIAMDeduplicatesWithinAStatement(t *testing.T) {
	t.Parallel()

	// Pike's real scan output repeats an action within a statement when
	// multiple attributes on one resource all require it (e.g.
	// aws_ecs_task_definition's execution_role_arn and task_role_arn
	// attributes both require ecs:DescribeTaskDefinition-adjacent calls).
	p := policy.Policy{Statements: []policy.Statement{{
		Actions:   []string{"ecs:DescribeTaskDefinition", "ecs:RegisterTaskDefinition", "ecs:DescribeTaskDefinition"},
		Resources: []string{"*", "*"},
	}}}

	doc, err := transform.DeployBoundary(p, testConfig())
	if err != nil {
		t.Fatalf("DeployBoundary() error = %v", err)
	}

	var nonIAM []string

	for _, s := range doc.Statements {
		if s.Sid == "NonIAM0" {
			nonIAM = s.Action
		}
	}

	want := []string{"ecs:DescribeTaskDefinition", "ecs:RegisterTaskDefinition"}
	if len(nonIAM) != len(want) {
		t.Fatalf("NonIAM0 Action = %v, want %v (deduplicated)", nonIAM, want)
	}

	for i, a := range want {
		if nonIAM[i] != a {
			t.Errorf("NonIAM0 Action[%d] = %q, want %q", i, nonIAM[i], a)
		}
	}
}
