// Package transform implements deploy-boundary.json generation (spec
// section 12): the deploy policy's requested permissions, rewritten so
// every IAM/STS write is either scoped to the managed-role path with the
// required conditions, or explicitly denied. Non-IAM permissions and
// read-only IAM/STS actions pass through unchanged; everything else in
// the transformation order is unconditional on what Pike actually
// detected — the mandatory denies exist to hold the invariants in
// spec section 4.2 regardless of the scanned Terraform.
package transform

import (
	"fmt"
	"sort"
	"strings"

	"github.com/jameswoolfenden/pike/internal/bootstrap/classify"
	"github.com/jameswoolfenden/pike/internal/bootstrap/config"
	"github.com/jameswoolfenden/pike/internal/bootstrap/render"
	"github.com/jameswoolfenden/pike/internal/policy"
)

// PassRoleServicesRequiredError is returned when the scanned policy
// requires iam:PassRole but pass_role.allowed_services is empty (spec
// section 12.3: "Generation fails if iam:PassRole is detected but no
// approved service is configured.").
type PassRoleServicesRequiredError struct{}

func (*PassRoleServicesRequiredError) Error() string {
	return "iam:PassRole detected but pass_role.allowed_services is empty"
}

// ServiceLinkedRoleServicesRequiredError is returned when the scanned
// policy requires iam:CreateServiceLinkedRole but
// service_linked_roles.allowed_services is empty (spec section 12.4).
type ServiceLinkedRoleServicesRequiredError struct{}

func (*ServiceLinkedRoleServicesRequiredError) Error() string {
	return "iam:CreateServiceLinkedRole detected but service_linked_roles.allowed_services is empty"
}

// UnknownIAMActionError is returned when the scanned policy requires an
// IAM/STS action the classifier doesn't recognize and
// security.fail_on_unknown_iam_action is true (the default). Spec
// section 4.2 rule 7: unknown IAM write actions fail generation by
// default rather than being silently allowed into the boundary.
type UnknownIAMActionError struct {
	Action string
}

func (e *UnknownIAMActionError) Error() string {
	return fmt.Sprintf("unclassified IAM/STS action %q: refusing to generate a boundary that silently allows it", e.Action)
}

// accessKeyActions and humanIdentityActions split classify.HumanPrincipal
// (one combined class for classification purposes) back into the two
// groups pike.yaml's security.deny_human_principal_management and
// security.deny_access_key_management flags each control independently.
var (
	humanIdentityActions = []string{
		"iam:CreateUser",
		"iam:DeleteUser",
		"iam:UpdateUser",
		"iam:CreateLoginProfile",
		"iam:UpdateLoginProfile",
		"iam:DeleteLoginProfile",
		"iam:CreateVirtualMFADevice",
		"iam:DeactivateMFADevice",
		"iam:DeleteVirtualMFADevice",
		"iam:EnableMFADevice",
		"iam:ResyncMFADevice",
	}

	accessKeyActions = []string{
		"iam:CreateAccessKey",
		"iam:UpdateAccessKey",
		"iam:DeleteAccessKey",
	}
)

// DeployBoundary transforms a scanned Policy into deploy-boundary.json:
// the maximum permissions the deploy role may hold once bootstrapped.
func DeployBoundary(p policy.Policy, cfg *config.Config) (render.Document, error) {
	grouped := groupByClass(p)

	if cfg.Security.FailOnUnknownIAMAction {
		for _, action := range grouped[classify.Unknown] {
			return render.Document{}, &UnknownIAMActionError{Action: action}
		}
	}

	var statements []render.Statement

	statements = append(statements, nonIAMStatements(p)...)

	if read := grouped[classify.Read]; len(read) > 0 {
		statements = append(statements, render.Statement{
			Sid:      "IAMRead",
			Effect:   "Allow",
			Action:   read,
			Resource: []string{"*"}, // no live account at generate time to scope further
		})
	}

	managedRoleARN := roleARN(cfg, cfg.ManagedRoles.Path)
	workloadBoundaryARN := namedPolicyARN(cfg, cfg.Policies.WorkloadBoundary.Path, cfg.Policies.WorkloadBoundary.Name)

	if len(grouped[classify.RoleCreate]) > 0 {
		statements = append(statements, render.Statement{
			Sid:      "CreateRoleWithRequiredBoundary",
			Effect:   "Allow",
			Action:   []string{"iam:CreateRole"},
			Resource: []string{managedRoleARN},
			Condition: &render.Condition{StringEquals: map[string]any{
				"iam:PermissionsBoundary": workloadBoundaryARN,
			}},
		})
	}

	rolePolicyWrite := grouped[classify.RolePolicyWrite]

	// iam:PutRolePermissionsBoundary is deliberately excluded from
	// ManageTerraformRoles (and given its own conditioned statement below)
	// even when Pike detects it directly — an unconditioned grant of this
	// action lets the deploy role re-point a role at a different, more
	// permissive boundary policy, silently bypassing the boundary the
	// CreateRole condition above and the boundary-tampering denies in
	// mandatoryDenies exist to enforce. See MaintainRequiredBoundary.
	if manageActions := removeAction(rolePolicyWrite, "iam:PutRolePermissionsBoundary"); len(manageActions) > 0 {
		statements = append(statements, render.Statement{
			Sid:      "ManageTerraformRoles",
			Effect:   "Allow",
			Action:   manageActions,
			Resource: []string{managedRoleARN},
		})
	}

	// The deploy role needs iam:PutRolePermissionsBoundary to maintain the
	// boundary on roles it manages (e.g. Terraform reapplying an existing
	// role's permissions_boundary attribute) whenever it creates or
	// modifies roles at all (spec section 12.2: "Actions required for
	// boundary lifecycle") — but only ever to (re)attach the one required
	// workload boundary, never anything else.
	if len(grouped[classify.RoleCreate]) > 0 || len(rolePolicyWrite) > 0 {
		statements = append(statements, render.Statement{
			Sid:      "MaintainRequiredBoundary",
			Effect:   "Allow",
			Action:   []string{"iam:PutRolePermissionsBoundary"},
			Resource: []string{managedRoleARN},
			Condition: &render.Condition{StringEquals: map[string]any{
				"iam:PermissionsBoundary": workloadBoundaryARN,
			}},
		})
	}

	if len(grouped[classify.PassRole]) > 0 {
		if len(cfg.PassRole.AllowedServices) == 0 {
			return render.Document{}, &PassRoleServicesRequiredError{}
		}

		statements = append(statements, render.Statement{
			Sid:      "PassManagedRolesToApprovedServices",
			Effect:   "Allow",
			Action:   []string{"iam:PassRole"},
			Resource: []string{managedRoleARN},
			Condition: &render.Condition{StringEquals: map[string]any{
				"iam:PassedToService": sortedCopy(cfg.PassRole.AllowedServices),
			}},
		})
	}

	if len(grouped[classify.ServiceLinkedRole]) > 0 {
		if len(cfg.ServiceLinkedRoles.AllowedServices) == 0 {
			return render.Document{}, &ServiceLinkedRoleServicesRequiredError{}
		}

		statements = append(statements, render.Statement{
			Sid:      "CreateApprovedServiceLinkedRoles",
			Effect:   "Allow",
			Action:   []string{"iam:CreateServiceLinkedRole"},
			Resource: []string{"*"},
			Condition: &render.Condition{StringEquals: map[string]any{
				"iam:AWSServiceName": sortedCopy(cfg.ServiceLinkedRoles.AllowedServices),
			}},
		})
	}

	statements = append(statements, mandatoryDenies(cfg, managedRoleARN)...)

	sort.Slice(statements, func(i, j int) bool { return statements[i].Sid < statements[j].Sid })

	return render.Document{Version: "2012-10-17", Statements: statements}, nil
}

// mandatoryDenies returns the always-present deny statements security's
// deny_* flags control (spec section 12.5). These are independent of
// what Pike detected — they exist to hold section 4.2's invariants
// (the deploy role can't remove or widen its own boundary, can't touch
// human principals, access keys or org governance) regardless of the
// Terraform being deployed.
func mandatoryDenies(cfg *config.Config, managedRoleARN string) []render.Statement {
	var denies []render.Statement

	if cfg.Security.DenyBoundaryRemoval {
		denies = append(denies, render.Statement{
			Sid:      "DenyWorkloadBoundaryRemoval",
			Effect:   "Deny",
			Action:   []string{"iam:DeleteRolePermissionsBoundary"},
			Resource: []string{managedRoleARN},
		})
	}

	if cfg.Security.DenyBoundaryPolicyMutation {
		// Covers both the deploy_boundary and workload_boundary policy
		// paths, not just deploy_boundary's — they're free to be
		// configured under different IAM paths (config.Validate only
		// requires them distinct from deploy_role/managed_roles, not from
		// each other), and the workload boundary is the one actually
		// constraining roles CreateRoleWithRequiredBoundary/
		// MaintainRequiredBoundary create. No statement currently grants
		// iam:CreatePolicyVersion et al. at all (PolicyMutation-classified
		// actions aren't wired into any Allow), so this is defense in
		// depth rather than closing an active hole - but it means a
		// future change that does grant one of these actions can't
		// silently reopen the boundary via its content instead of its
		// attachment.
		denies = append(denies, render.Statement{
			Sid:    "DenyBoundaryPolicyMutation",
			Effect: "Deny",
			Action: sortedCopy([]string{
				"iam:CreatePolicyVersion",
				"iam:DeletePolicyVersion",
				"iam:SetDefaultPolicyVersion",
				"iam:DeletePolicy",
			}),
			Resource: sortedUnique([]string{
				policyPathARN(cfg, cfg.Policies.DeployBoundary.Path),
				policyPathARN(cfg, cfg.Policies.WorkloadBoundary.Path),
			}),
		})
	}

	if cfg.Security.DenyHumanPrincipalManagement {
		denies = append(denies, render.Statement{
			Sid:      "DenyHumanPrincipalManagement",
			Effect:   "Deny",
			Action:   sortedCopy(humanIdentityActions),
			Resource: []string{"*"},
		})
	}

	if cfg.Security.DenyAccessKeyManagement {
		denies = append(denies, render.Statement{
			Sid:      "DenyAccessKeyManagement",
			Effect:   "Deny",
			Action:   sortedCopy(accessKeyActions),
			Resource: []string{"*"},
		})
	}

	if cfg.Security.DenyOrganizationManagement {
		denies = append(denies, render.Statement{
			Sid:      "DenyOrganizationManagement",
			Effect:   "Deny",
			Action:   []string{"account:*", "organizations:*"},
			Resource: []string{"*"},
		})
	}

	return denies
}

// nonIAMStatements copies every statement of the scanned policy that
// contains no IAM/STS actions through unchanged (transformation order
// step 1) — the boundary must allow at least what non-IAM services need,
// since it only restricts IAM.
func nonIAMStatements(p policy.Policy) []render.Statement {
	var statements []render.Statement

	for i, s := range p.Statements {
		var nonIAM []string

		for _, action := range s.Actions {
			if !isIAMOrSTS(action) {
				nonIAM = append(nonIAM, action)
			}
		}

		if len(nonIAM) == 0 {
			continue
		}

		statements = append(statements, render.Statement{
			Sid:      fmt.Sprintf("NonIAM%d", i),
			Effect:   "Allow",
			Action:   sortedUnique(nonIAM),
			Resource: sortedUnique(s.Resources),
		})
	}

	return statements
}

// groupByClass classifies every distinct IAM/STS action in p and groups
// them by class, sorted within each group. Non-IAM/STS actions are
// excluded entirely — they're handled by nonIAMStatements instead.
func groupByClass(p policy.Policy) map[classify.Class][]string {
	grouped := make(map[classify.Class][]string)
	seen := make(map[string]bool)

	for _, action := range p.Actions() {
		if !isIAMOrSTS(action) || seen[action] {
			continue
		}

		seen[action] = true
		class := classify.Classify(action)
		grouped[class] = append(grouped[class], action)
	}

	for class := range grouped {
		sort.Strings(grouped[class])
	}

	return grouped
}

func isIAMOrSTS(action string) bool {
	return strings.HasPrefix(action, "iam:") || strings.HasPrefix(action, "sts:")
}

func roleARN(cfg *config.Config, path string) string {
	return fmt.Sprintf("arn:%s:iam::%s:role%s*", cfg.AWS.Partition, cfg.AWS.AccountID, path)
}

func namedPolicyARN(cfg *config.Config, path, name string) string {
	return fmt.Sprintf("arn:%s:iam::%s:policy%s%s", cfg.AWS.Partition, cfg.AWS.AccountID, path, name)
}

func policyPathARN(cfg *config.Config, path string) string {
	return fmt.Sprintf("arn:%s:iam::%s:policy%s*", cfg.AWS.Partition, cfg.AWS.AccountID, path)
}

func sortedCopy(s []string) []string {
	out := append([]string(nil), s...)
	sort.Strings(out)

	return out
}

// sortedUnique returns the sorted, deduplicated contents of s. Pike's
// scan output can list the same action more than once within a
// statement, and a rendered policy document shouldn't repeat it.
func sortedUnique(s []string) []string {
	seen := make(map[string]bool, len(s))

	var out []string

	for _, v := range s {
		if seen[v] {
			continue
		}

		seen[v] = true

		out = append(out, v)
	}

	sort.Strings(out)

	return out
}

// removeAction returns a sorted copy of actions with target removed, if
// present. Used to keep iam:PutRolePermissionsBoundary out of
// ManageTerraformRoles regardless of whether it arrived there via direct
// detection — it always gets its own conditioned statement instead (see
// MaintainRequiredBoundary in DeployBoundary).
func removeAction(actions []string, target string) []string {
	var out []string

	for _, a := range actions {
		if a != target {
			out = append(out, a)
		}
	}

	sort.Strings(out)

	return out
}
