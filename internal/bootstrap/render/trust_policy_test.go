package render_test

import (
	"errors"
	"testing"

	"github.com/jameswoolfenden/pike/internal/bootstrap/config"
	"github.com/jameswoolfenden/pike/internal/bootstrap/render"
)

func testConfig() *config.Config {
	return &config.Config{
		AWS: config.AWSConfig{AccountID: "123456789012", Partition: "aws"},
		Principal: config.PrincipalConfig{
			Type: "iam-user",
			Name: "cli-user",
			MFA:  config.MFAConfig{Required: true, MaxAgeSeconds: 3600},
		},
		DeployRole: config.DeployRoleConfig{Name: "TerraformDeployRole", Path: "/terraform-deploy/"},
	}
}

func TestTrustPolicy(t *testing.T) {
	t.Parallel()

	doc, err := render.TrustPolicy(testConfig())
	if err != nil {
		t.Fatalf("TrustPolicy() error = %v", err)
	}

	if len(doc.Statement) != 1 {
		t.Fatalf("len(Statement) = %d, want 1", len(doc.Statement))
	}

	s := doc.Statement[0]

	if s.Principal["AWS"] != "arn:aws:iam::123456789012:user/cli-user" {
		t.Errorf("Principal[AWS] = %q, want the user ARN", s.Principal["AWS"])
	}

	if s.Action != "sts:AssumeRole" {
		t.Errorf("Action = %q, want sts:AssumeRole", s.Action)
	}

	if s.Condition == nil {
		t.Fatal("Condition = nil, want MFA condition since principal.mfa.required is true")
	}

	if s.Condition.Bool["aws:MultiFactorAuthPresent"] != "true" {
		t.Errorf("Condition.Bool[aws:MultiFactorAuthPresent] = %q, want \"true\"", s.Condition.Bool["aws:MultiFactorAuthPresent"])
	}

	if s.Condition.NumericLessThanEquals["aws:MultiFactorAuthAge"] != "3600" {
		t.Errorf("Condition.NumericLessThanEquals[aws:MultiFactorAuthAge] = %q, want \"3600\"", s.Condition.NumericLessThanEquals["aws:MultiFactorAuthAge"])
	}
}

func TestTrustPolicy_MFANotRequired(t *testing.T) {
	t.Parallel()

	cfg := testConfig()
	cfg.Principal.MFA.Required = false

	doc, err := render.TrustPolicy(cfg)
	if err != nil {
		t.Fatalf("TrustPolicy() error = %v", err)
	}

	if doc.Statement[0].Condition != nil {
		t.Error("Condition present despite principal.mfa.required = false")
	}
}

func TestTrustPolicy_UnsupportedPrincipalType(t *testing.T) {
	t.Parallel()

	cfg := testConfig()
	cfg.Principal.Type = "identity-center"

	_, err := render.TrustPolicy(cfg)

	var target *render.UnsupportedPrincipalTypeError
	if !errors.As(err, &target) {
		t.Fatalf("TrustPolicy() error = %v, want *UnsupportedPrincipalTypeError", err)
	}
}

func TestAssumeRolePolicy(t *testing.T) {
	t.Parallel()

	doc := render.AssumeRolePolicy(testConfig())

	if len(doc.Statements) != 1 {
		t.Fatalf("len(Statements) = %d, want 1", len(doc.Statements))
	}

	s := doc.Statements[0]

	if len(s.Action) != 1 || s.Action[0] != "sts:AssumeRole" {
		t.Errorf("Action = %v, want [sts:AssumeRole]", s.Action)
	}

	if len(s.Resource) != 1 || s.Resource[0] != "arn:aws:iam::123456789012:role/terraform-deploy/TerraformDeployRole" {
		t.Errorf("Resource = %v, want the deploy role ARN", s.Resource)
	}

	if s.Condition != nil {
		t.Error("Condition present — the source-principal policy must grant only sts:AssumeRole, nothing conditional")
	}
}
