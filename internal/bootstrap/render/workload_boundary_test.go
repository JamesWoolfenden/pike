package render_test

import (
	"errors"
	"testing"

	"github.com/jameswoolfenden/pike/internal/bootstrap/config"
	"github.com/jameswoolfenden/pike/internal/bootstrap/render"
)

func TestWorkloadBoundary(t *testing.T) {
	t.Parallel()

	cfg := &config.Config{
		Policies: config.PoliciesConfig{
			WorkloadBoundary: config.WorkloadBoundaryConfig{
				Mode:           "allowlist",
				AllowedActions: []string{"logs:PutLogEvents", "ecr:GetAuthorizationToken"},
			},
		},
	}

	doc, err := render.WorkloadBoundary(cfg)
	if err != nil {
		t.Fatalf("WorkloadBoundary() error = %v", err)
	}

	if len(doc.Statements) != 1 {
		t.Fatalf("len(Statements) = %d, want 1", len(doc.Statements))
	}

	s := doc.Statements[0]
	if s.Effect != "Allow" {
		t.Errorf("Effect = %q, want Allow", s.Effect)
	}

	want := []string{"ecr:GetAuthorizationToken", "logs:PutLogEvents"}
	if len(s.Action) != len(want) || s.Action[0] != want[0] || s.Action[1] != want[1] {
		t.Errorf("Action = %v, want sorted %v", s.Action, want)
	}

	if len(s.Resource) != 1 || s.Resource[0] != "*" {
		t.Errorf("Resource = %v, want [*]", s.Resource)
	}
}

func TestWorkloadBoundary_UnsupportedMode(t *testing.T) {
	t.Parallel()

	cfg := &config.Config{
		Policies: config.PoliciesConfig{
			WorkloadBoundary: config.WorkloadBoundaryConfig{Mode: "infer"},
		},
	}

	_, err := render.WorkloadBoundary(cfg)

	var target *render.UnsupportedWorkloadBoundaryModeError
	if !errors.As(err, &target) {
		t.Fatalf("WorkloadBoundary() error = %v, want *UnsupportedWorkloadBoundaryModeError", err)
	}
}
