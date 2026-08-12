package transform_test

import (
	"regexp"
	"slices"
	"strings"
	"testing"

	"github.com/jameswoolfenden/pike/internal/bootstrap/render"
	"github.com/jameswoolfenden/pike/internal/bootstrap/transform"
	"github.com/jameswoolfenden/pike/internal/policy"
)

// This file simulates specific attempted IAM privilege-escalation calls
// against a generated deploy-boundary.json and asserts the whole document
// (Allow and Deny statements together, with explicit-deny-wins evaluation)
// actually blocks them - rather than checking that a particular action
// string is merely present or absent in one statement, which is what let
// the original iam:PutRolePermissionsBoundary bug go unnoticed.

// simulatedRequest is a minimal stand-in for an AWS API call: the action
// and resource being requested, and the condition-key/value context that
// would be present on the real request (e.g. iam:PermissionsBoundary is
// only present on calls that set a boundary).
type simulatedRequest struct {
	action   string
	resource string
	context  map[string]string
}

// evaluateAllowed mimics AWS IAM policy evaluation order: an explicit Deny
// in any applicable statement always wins, regardless of any Allow
// elsewhere in the document; otherwise the request is allowed only if at
// least one Allow statement applies; everything else is an implicit deny.
func evaluateAllowed(doc render.Document, req simulatedRequest) bool {
	allowed := false

	for _, s := range doc.Statements {
		if !statementApplies(s, req) {
			continue
		}

		if s.Effect == "Deny" {
			return false
		}

		allowed = true
	}

	return allowed
}

func statementApplies(s render.Statement, req simulatedRequest) bool {
	if !containsAction(s.Action, req.action) {
		return false
	}

	if !matchesAny(s.Resource, req.resource) {
		return false
	}

	return conditionSatisfied(s.Condition, req.context)
}

func containsAction(actions []string, action string) bool {
	return slices.Contains(actions, action)
}

func matchesAny(patterns []string, resource string) bool {
	return slices.ContainsFunc(patterns, func(p string) bool {
		return matchesARNPattern(p, resource)
	})
}

// matchesARNPattern treats "*" as a glob wildcard, the only wildcard usage
// this generator ever produces (trailing-* role/policy paths, or a bare
// "*").
func matchesARNPattern(pattern, resource string) bool {
	quoted := regexp.QuoteMeta(pattern)
	re := "^" + strings.ReplaceAll(quoted, `\*`, ".*") + "$"

	matched, err := regexp.MatchString(re, resource)
	if err != nil {
		return false
	}

	return matched
}

// conditionSatisfied evaluates a StringEquals condition against the
// request context the way AWS does: every key in the condition must be
// present in the request context and match, and a nil Condition always
// matches (the statement is unconditional).
func conditionSatisfied(cond *render.Condition, ctx map[string]string) bool {
	if cond == nil || cond.StringEquals == nil {
		return true
	}

	for key, want := range cond.StringEquals {
		got, present := ctx[key]
		if !present {
			return false
		}

		if !valueMatches(want, got) {
			return false
		}
	}

	return true
}

func valueMatches(want any, got string) bool {
	switch v := want.(type) {
	case string:
		return v == got
	case []string:
		return slices.Contains(v, got)
	default:
		return false
	}
}

const (
	managedRoleWildcard           = "arn:aws:iam::123456789012:role/terraform-managed/*"
	someManagedRole               = "arn:aws:iam::123456789012:role/terraform-managed/some-app-role"
	requiredWorkloadBoundaryARN   = "arn:aws:iam::123456789012:policy/boundaries/TerraformWorkloadBoundary"
	attackerControlledBoundaryARN = "arn:aws:iam::123456789012:policy/boundaries/AttackerControlledPermissive"
)

// TestPrivesc_BoundarySwapBlocked simulates the exact escalation the
// original bug allowed: the deploy role calling PutRolePermissionsBoundary
// on a role it manages, but pointing it at a different, attacker-controlled
// boundary policy instead of the required workload boundary. A generated
// policy that lets this succeed defeats the entire purpose of
// CreateRoleWithRequiredBoundary - the created role's boundary can just be
// swapped out afterwards.
func TestPrivesc_BoundarySwapBlocked(t *testing.T) {
	t.Parallel()

	p := policy.Policy{Statements: []policy.Statement{{Actions: []string{"iam:CreateRole"}, Resources: []string{"*"}}}}

	doc, err := transform.DeployBoundary(p, testConfig())
	if err != nil {
		t.Fatalf("DeployBoundary() error = %v", err)
	}

	req := simulatedRequest{
		action:   "iam:PutRolePermissionsBoundary",
		resource: someManagedRole,
		context:  map[string]string{"iam:PermissionsBoundary": attackerControlledBoundaryARN},
	}

	if evaluateAllowed(doc, req) {
		t.Fatal("boundary-swap escalation succeeded: the deploy role could re-point a managed role at an attacker-controlled permissions boundary")
	}
}

// TestPrivesc_BoundarySwap_MethodologyCatchesTheOriginalBug is a
// calibration test: it hand-builds the exact statement shape the original
// (pre-fix) fork produced - PutRolePermissionsBoundary bundled into
// ManageTerraformRoles with no Condition - and confirms evaluateAllowed
// correctly flags it as exploitable. This proves the simulator above isn't
// vacuously passing; it would have caught the real bug had it existed when
// the bug did.
func TestPrivesc_BoundarySwap_MethodologyCatchesTheOriginalBug(t *testing.T) {
	t.Parallel()

	vulnerableDoc := render.Document{
		Version: "2012-10-17",
		Statements: []render.Statement{
			{
				Sid:      "ManageTerraformRoles",
				Effect:   "Allow",
				Action:   []string{"iam:PutRolePermissionsBoundary", "iam:PutRolePolicy"},
				Resource: []string{managedRoleWildcard},
				// No Condition - this is the original bug's shape.
			},
		},
	}

	req := simulatedRequest{
		action:   "iam:PutRolePermissionsBoundary",
		resource: someManagedRole,
		context:  map[string]string{"iam:PermissionsBoundary": attackerControlledBoundaryARN},
	}

	if !evaluateAllowed(vulnerableDoc, req) {
		t.Fatal("evaluateAllowed did not flag the known-vulnerable unconditioned statement shape as exploitable - the simulator itself is broken")
	}
}

// TestPrivesc_LegitimateBoundaryMaintenanceStillWorks confirms the fix
// didn't just block everything: reapplying the *required* boundary (e.g.
// Terraform correcting drift on an existing managed role) must still
// succeed.
func TestPrivesc_LegitimateBoundaryMaintenanceStillWorks(t *testing.T) {
	t.Parallel()

	p := policy.Policy{Statements: []policy.Statement{{Actions: []string{"iam:CreateRole"}, Resources: []string{"*"}}}}

	doc, err := transform.DeployBoundary(p, testConfig())
	if err != nil {
		t.Fatalf("DeployBoundary() error = %v", err)
	}

	req := simulatedRequest{
		action:   "iam:PutRolePermissionsBoundary",
		resource: someManagedRole,
		context:  map[string]string{"iam:PermissionsBoundary": requiredWorkloadBoundaryARN},
	}

	if !evaluateAllowed(doc, req) {
		t.Fatal("legitimate boundary maintenance (reapplying the required workload boundary) was blocked - the fix over-corrected")
	}
}

// TestPrivesc_PassRoleToUnapprovedServiceBlocked simulates the classic
// PassRole escalation: pass a managed role to a service the deploy role
// can also provision compute in (e.g. Lambda), then run code as that role.
// pass_role.allowed_services in testConfig() only allows ECS services, so
// passing to Lambda must fail.
func TestPrivesc_PassRoleToUnapprovedServiceBlocked(t *testing.T) {
	t.Parallel()

	p := policy.Policy{Statements: []policy.Statement{{Actions: []string{"iam:PassRole"}, Resources: []string{"*"}}}}

	doc, err := transform.DeployBoundary(p, testConfig())
	if err != nil {
		t.Fatalf("DeployBoundary() error = %v", err)
	}

	req := simulatedRequest{
		action:   "iam:PassRole",
		resource: someManagedRole,
		context:  map[string]string{"iam:PassedToService": "lambda.amazonaws.com"},
	}

	if evaluateAllowed(doc, req) {
		t.Fatal("iam:PassRole succeeded against an unapproved service (lambda.amazonaws.com) - pass_role.allowed_services was not enforced")
	}
}

// TestPrivesc_PassRoleToApprovedServiceStillWorks is the sanity check for
// the test above: passing to a service that *is* on the allow-list must
// still succeed.
func TestPrivesc_PassRoleToApprovedServiceStillWorks(t *testing.T) {
	t.Parallel()

	p := policy.Policy{Statements: []policy.Statement{{Actions: []string{"iam:PassRole"}, Resources: []string{"*"}}}}

	doc, err := transform.DeployBoundary(p, testConfig())
	if err != nil {
		t.Fatalf("DeployBoundary() error = %v", err)
	}

	req := simulatedRequest{
		action:   "iam:PassRole",
		resource: someManagedRole,
		context:  map[string]string{"iam:PassedToService": "ecs-tasks.amazonaws.com"},
	}

	if !evaluateAllowed(doc, req) {
		t.Fatal("iam:PassRole was blocked even for an approved service (ecs-tasks.amazonaws.com)")
	}
}

// TestPrivesc_BoundaryPolicyContentMutationDeniedAcrossBothPaths locks in
// the DenyBoundaryPolicyMutation hardening: deploy_boundary and
// workload_boundary are allowed to live under different IAM paths
// (config.Validate doesn't require them to match each other), and the
// deny must cover both - not just deploy_boundary's, which was the
// original scope. No current statement actually grants
// iam:CreatePolicyVersion (PolicyMutation-classified actions are never
// wired into an Allow), so this test injects a hypothetical Allow to prove
// the Deny would still win via explicit-deny-always-wins evaluation if
// that ever changed.
func TestPrivesc_BoundaryPolicyContentMutationDeniedAcrossBothPaths(t *testing.T) {
	t.Parallel()

	cfg := testConfig()
	cfg.Policies.DeployBoundary.Path = "/deploy-boundaries/"
	cfg.Policies.WorkloadBoundary.Path = "/workload-boundaries/"

	doc, err := transform.DeployBoundary(policy.Policy{}, cfg)
	if err != nil {
		t.Fatalf("DeployBoundary() error = %v", err)
	}

	workloadBoundaryPolicyARN := "arn:aws:iam::123456789012:policy/workload-boundaries/TerraformWorkloadBoundary"

	doc.Statements = append(doc.Statements, render.Statement{
		Sid:      "HypotheticalGrant",
		Effect:   "Allow",
		Action:   []string{"iam:CreatePolicyVersion"},
		Resource: []string{workloadBoundaryPolicyARN},
	})

	req := simulatedRequest{
		action:   "iam:CreatePolicyVersion",
		resource: workloadBoundaryPolicyARN,
	}

	if evaluateAllowed(doc, req) {
		t.Fatal("iam:CreatePolicyVersion on the workload boundary policy succeeded - DenyBoundaryPolicyMutation doesn't cover workload_boundary's path")
	}
}
