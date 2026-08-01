package render_test

import (
	"encoding/json"
	"testing"

	"github.com/jameswoolfenden/pike/internal/bootstrap/render"
	"github.com/jameswoolfenden/pike/internal/policy"
)

func TestDeployPolicy(t *testing.T) {
	t.Parallel()

	p := policy.Policy{
		Statements: []policy.Statement{
			{Actions: []string{"iam:CreateRole", "iam:GetRole"}, Resources: []string{"*"}},
			{Actions: []string{"ec2:RunInstances", "ec2:DescribeInstances"}, Resources: []string{"*"}},
		},
	}

	doc := render.DeployPolicy(p)

	if doc.Version != "2012-10-17" {
		t.Errorf("Version = %q, want 2012-10-17", doc.Version)
	}

	if len(doc.Statements) != 2 {
		t.Fatalf("len(Statements) = %d, want 2", len(doc.Statements))
	}

	// Statements ordered by first action: "ec2:..." sorts before "iam:...".
	if got := doc.Statements[0].Action; len(got) != 2 || got[0] != "ec2:DescribeInstances" || got[1] != "ec2:RunInstances" {
		t.Errorf("Statements[0].Action = %v, want sorted [ec2:DescribeInstances ec2:RunInstances]", got)
	}

	if got := doc.Statements[1].Action; len(got) != 2 || got[0] != "iam:CreateRole" || got[1] != "iam:GetRole" {
		t.Errorf("Statements[1].Action = %v, want sorted [iam:CreateRole iam:GetRole]", got)
	}

	for _, s := range doc.Statements {
		if s.Effect != "Allow" {
			t.Errorf("Statement.Effect = %q, want Allow (deploy-policy carries no deny semantics)", s.Effect)
		}
	}
}

func TestDeployPolicy_DropsEmptyStatements(t *testing.T) {
	t.Parallel()

	p := policy.Policy{
		Statements: []policy.Statement{
			{Actions: nil, Resources: []string{"*"}},
			{Actions: []string{"iam:GetRole"}, Resources: []string{"*"}},
		},
	}

	doc := render.DeployPolicy(p)

	if len(doc.Statements) != 1 {
		t.Fatalf("len(Statements) = %d, want 1 (the actionless statement should be dropped)", len(doc.Statements))
	}
}

func TestDeployPolicy_DeduplicatesWithinAStatement(t *testing.T) {
	t.Parallel()

	// Pike's real scan output repeats an action within a statement when
	// multiple attributes on one resource all require it (e.g. several
	// attributes on aws_iam_role all requiring iam:GetRole).
	p := policy.Policy{
		Statements: []policy.Statement{
			{Actions: []string{"iam:GetRole", "iam:PassRole", "iam:GetRole", "iam:PassRole"}, Resources: []string{"*", "*"}},
		},
	}

	doc := render.DeployPolicy(p)

	if len(doc.Statements) != 1 {
		t.Fatalf("len(Statements) = %d, want 1", len(doc.Statements))
	}

	got := doc.Statements[0].Action
	want := []string{"iam:GetRole", "iam:PassRole"}

	if len(got) != len(want) {
		t.Fatalf("Action = %v, want %v (deduplicated)", got, want)
	}

	for i, a := range want {
		if got[i] != a {
			t.Errorf("Action[%d] = %q, want %q", i, got[i], a)
		}
	}
}

func TestDeployPolicy_Deterministic(t *testing.T) {
	t.Parallel()

	p := policy.Policy{
		Statements: []policy.Statement{
			{Actions: []string{"s3:PutObject", "s3:GetObject"}, Resources: []string{"*"}},
			{Actions: []string{"iam:CreateRole"}, Resources: []string{"*"}},
		},
	}

	first, err := json.Marshal(render.DeployPolicy(p))
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		got, err := json.Marshal(render.DeployPolicy(p))
		if err != nil {
			t.Fatal(err)
		}

		if string(got) != string(first) {
			t.Fatalf("run %d: DeployPolicy() rendered differently:\n%s\nvs\n%s", i, got, first)
		}
	}
}

func TestEscalationActions(t *testing.T) {
	t.Parallel()

	p := policy.Policy{
		Statements: []policy.Statement{
			{Actions: []string{
				"iam:GetRole",               // read — excluded
				"iam:CreateRole",            // escalation-sensitive
				"iam:PassRole",              // escalation-sensitive
				"ec2:DescribeInstances",     // not IAM/STS at all — excluded
				"ec2:RunInstances",          // not IAM/STS at all — excluded
				"sts:GetCallerIdentity",     // read — excluded
				"iam:UpdateRoleDescription", // unmapped IAM write — still escalation-sensitive
			}},
		},
	}

	got := render.EscalationActions(p)

	want := []string{"iam:CreateRole", "iam:PassRole", "iam:UpdateRoleDescription"}
	if len(got) != len(want) {
		t.Fatalf("EscalationActions() = %v, want %v", got, want)
	}

	for i, action := range want {
		if got[i] != action {
			t.Errorf("EscalationActions()[%d] = %q, want %q", i, got[i], action)
		}
	}
}

func TestEscalationActions_Empty(t *testing.T) {
	t.Parallel()

	p := policy.Policy{Statements: []policy.Statement{{Actions: []string{"ec2:DescribeInstances", "iam:GetRole"}}}}

	if got := render.EscalationActions(p); len(got) != 0 {
		t.Errorf("EscalationActions() = %v, want empty", got)
	}
}
