package render

import (
	"fmt"

	"github.com/jameswoolfenden/pike/internal/bootstrap/config"
)

// UnsupportedWorkloadBoundaryModeError is returned when
// policies.workload_boundary.mode isn't one v0.1 knows how to render.
// Only "allowlist" is implemented; "infer" is spec section 13.3's future
// mode, not yet built.
type UnsupportedWorkloadBoundaryModeError struct {
	Mode string
}

func (e *UnsupportedWorkloadBoundaryModeError) Error() string {
	return fmt.Sprintf("unsupported workload_boundary.mode %q (only \"allowlist\" is implemented)", e.Mode)
}

// WorkloadBoundary renders workload-boundary.json: the maximum
// permissions for Terraform-created workload roles. Per spec section
// 13.1, this comes entirely from explicit configuration — it must not be
// derived from the deploy policy, so unlike DeployPolicy/DeployBoundary
// this takes no scanned Policy at all.
func WorkloadBoundary(cfg *config.Config) (Document, error) {
	if cfg.Policies.WorkloadBoundary.Mode != "allowlist" {
		return Document{}, &UnsupportedWorkloadBoundaryModeError{Mode: cfg.Policies.WorkloadBoundary.Mode}
	}

	return Document{
		Version: policyVersion,
		Statements: []Statement{{
			Sid:      "TerraformWorkloadBoundary",
			Effect:   "Allow",
			Action:   sortedUnique(cfg.Policies.WorkloadBoundary.AllowedActions),
			Resource: []string{"*"},
		}},
	}, nil
}
