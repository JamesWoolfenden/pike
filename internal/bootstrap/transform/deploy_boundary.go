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

	"github.com/qj0r9j0vc2/rampart/internal/bootstrap/classify"
	"github.com/qj0r9j0vc2/rampart/internal/bootstrap/config"
	"github.com/qj0r9j0vc2/rampart/internal/bootstrap/render"
	"github.com/qj0r9j0vc2/rampart/internal/policy"
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
// groups rampart.yaml's security.deny_human_principal_management and
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

// managedRoleLifecycleActions are always included in ManageTerraformRoles
// when it's generated at all, regardless of what Pike detected — the
// deploy role needs them to fulfil the boundary lifecycle itself (spec
// section 12.2: "Actions required for boundary lifecycle").
var managedRoleLifecycleActions = []string{"iam:PutRolePermissionsBoundary"}

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

	if len(grouped[classify.RoleCreate]) > 0 {
		statements = append(statements, render.Statement{
			Sid:      "CreateRoleWithRequiredBoundary",
			Effect:   "Allow",
			Action:   []string{"iam:CreateRole"},
			Resource: []string{managedRoleARN},
			Condition: &render.Condition{StringEquals: map[string]any{
				"iam:PermissionsBoundary": namedPolicyARN(cfg, cfg.Policies.WorkloadBoundary.Path, cfg.Policies.WorkloadBoundary.Name),
			}},
		})
	}

	if manageActions := union(grouped[classify.RolePolicyWrite], managedRoleLifecycleActions); len(grouped[classify.RoleCreate]) > 0 || len(grouped[classify.RolePolicyWrite]) > 0 {
		statements = append(statements, render.Statement{
			Sid:      "ManageTerraformRoles",
			Effect:   "Allow",
			Action:   manageActions,
			Resource: []string{managedRoleARN},
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
		denies = append(denies, render.Statement{
			Sid:    "DenyBoundaryPolicyMutation",
			Effect: "Deny",
			Action: sortedCopy([]string{
				"iam:CreatePolicyVersion",
				"iam:DeletePolicyVersion",
				"iam:SetDefaultPolicyVersion",
				"iam:DeletePolicy",
			}),
			Resource: []string{policyPathARN(cfg, cfg.Policies.DeployBoundary.Path)},
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

// union returns the sorted, deduplicated union of a and b.
func union(a, b []string) []string {
	seen := make(map[string]bool, len(a)+len(b))

	var out []string

	for _, s := range append(append([]string(nil), a...), b...) {
		if seen[s] {
			continue
		}

		seen[s] = true

		out = append(out, s)
	}

	sort.Strings(out)

	return out
}
