// Package validate checks generated bootstrap artifacts against the
// blocking and warning rules in spec section 17. It runs as a defense-in-
// depth pass over the rendered documents themselves — deploy-boundary
// transformation (internal/bootstrap/transform) already enforces most of
// these invariants by construction, but validating the artifact
// independently also catches artifacts that were hand-edited after
// generation, and is what the future `pike bootstrap validate` CLI
// command will run against files already on disk.
//
// RMP-E010, RMP-E011 and RMP-E012 are not implemented here: they need
// live AWS state (a managed-path role audit, the caller's STS identity)
// that only exists once the inspect/apply milestone (M3/M4) is built.
// Likewise, most of section 17.2's warnings need infrastructure this
// package doesn't have yet (Pike mapping metadata, existing live AWS
// policy content) — only the workload-allowlist-wildcard warning is
// checkable from config alone, so it's the only one implemented so far.
package validate

import (
	"fmt"
	"strings"

	"github.com/jameswoolfenden/pike/internal/bootstrap/classify"
	"github.com/jameswoolfenden/pike/internal/bootstrap/config"
	"github.com/jameswoolfenden/pike/internal/bootstrap/render"
)

// Severity distinguishes findings that must block generation from ones
// that are informational only (spec sections 17.1 vs 17.2).
type Severity string

const (
	Blocking Severity = "blocking"
	Warning  Severity = "warning"
)

// Finding is a single validation result.
type Finding struct {
	Code     string
	Severity Severity
	Message  string
}

// Artifacts bundles the generated documents this package's rules check.
type Artifacts struct {
	DeployBoundary   render.Document
	TrustPolicy      render.TrustDocument
	AssumeRolePolicy render.Document
}

// Validate runs every rule in this package against a generated bootstrap
// package and its config, returning every finding (not just the first).
func Validate(a Artifacts, cfg *config.Config) []Finding {
	var findings []Finding

	findings = append(findings, checkCreateRoleConditioned(a.DeployBoundary, cfg)...)
	findings = append(findings, checkPassRole(a.DeployBoundary, cfg)...)
	findings = append(findings, checkServiceLinkedRole(a.DeployBoundary)...)
	findings = append(findings, checkBoundaryDenies(a.DeployBoundary)...)
	findings = append(findings, checkUnclassifiedActions(a.DeployBoundary)...)
	findings = append(findings, checkTrustPolicyMFA(a.TrustPolicy)...)
	findings = append(findings, checkAssumeRolePolicyMinimal(a.AssumeRolePolicy)...)
	findings = append(findings, checkWorkloadAllowlistWildcard(cfg)...)

	return findings
}

// checkCreateRoleConditioned covers RMP-E001 (unconditional CreateRole)
// and RMP-E002 (condition present but not the exact expected boundary).
func checkCreateRoleConditioned(doc render.Document, cfg *config.Config) []Finding {
	var findings []Finding

	expectedBoundaryARN := namedPolicyARN(cfg, cfg.Policies.WorkloadBoundary.Path, cfg.Policies.WorkloadBoundary.Name)

	for _, s := range doc.Statements {
		if s.Effect != "Allow" || !containsAction(s.Action, "iam:CreateRole") {
			continue
		}

		if s.Condition == nil {
			findings = append(findings, Finding{
				Code:     "RMP-E001",
				Severity: Blocking,
				Message:  "iam:CreateRole is allowed without any Condition",
			})

			continue
		}

		got := conditionString(s.Condition, "iam:PermissionsBoundary")
		if got != expectedBoundaryARN {
			findings = append(findings, Finding{
				Code:     "RMP-E002",
				Severity: Blocking,
				Message:  fmt.Sprintf("iam:CreateRole Condition.StringEquals[iam:PermissionsBoundary] = %q, want %q", got, expectedBoundaryARN),
			})
		}
	}

	return findings
}

// checkPassRole covers RMP-E003 (Resource "*" or outside the managed-role
// path) and RMP-E004 (missing PassedToService restriction).
func checkPassRole(doc render.Document, cfg *config.Config) []Finding {
	var findings []Finding

	managedPrefix := strings.TrimSuffix(roleARN(cfg, cfg.ManagedRoles.Path), "*")

	for _, s := range doc.Statements {
		if s.Effect != "Allow" || !containsAction(s.Action, "iam:PassRole") {
			continue
		}

		for _, r := range s.Resource {
			if r == "*" || !strings.HasPrefix(r, managedPrefix) {
				findings = append(findings, Finding{
					Code:     "RMP-E003",
					Severity: Blocking,
					Message:  fmt.Sprintf("iam:PassRole Resource %q is \"*\" or escapes the managed-role path", r),
				})
			}
		}

		if len(conditionStringList(s.Condition, "iam:PassedToService")) == 0 {
			findings = append(findings, Finding{
				Code:     "RMP-E004",
				Severity: Blocking,
				Message:  "iam:PassRole is missing a Condition.StringEquals[iam:PassedToService] restriction",
			})
		}
	}

	return findings
}

// checkServiceLinkedRole covers RMP-E005.
func checkServiceLinkedRole(doc render.Document) []Finding {
	var findings []Finding

	for _, s := range doc.Statements {
		if s.Effect != "Allow" || !containsAction(s.Action, "iam:CreateServiceLinkedRole") {
			continue
		}

		if len(conditionStringList(s.Condition, "iam:AWSServiceName")) == 0 {
			findings = append(findings, Finding{
				Code:     "RMP-E005",
				Severity: Blocking,
				Message:  "iam:CreateServiceLinkedRole is missing a Condition.StringEquals[iam:AWSServiceName] restriction",
			})
		}
	}

	return findings
}

// checkBoundaryDenies covers RMP-E006: boundary removal and boundary-
// policy mutation must each have an explicit Deny.
func checkBoundaryDenies(doc render.Document) []Finding {
	var findings []Finding

	if !hasDenyFor(doc, "iam:DeleteRolePermissionsBoundary") {
		findings = append(findings, Finding{
			Code:     "RMP-E006",
			Severity: Blocking,
			Message:  "no explicit Deny for iam:DeleteRolePermissionsBoundary (boundary removal)",
		})
	}

	for _, action := range []string{"iam:CreatePolicyVersion", "iam:DeletePolicyVersion", "iam:SetDefaultPolicyVersion", "iam:DeletePolicy"} {
		if !hasDenyFor(doc, action) {
			findings = append(findings, Finding{
				Code:     "RMP-E006",
				Severity: Blocking,
				Message:  fmt.Sprintf("no explicit Deny for %s (boundary-policy mutation)", action),
			})
		}
	}

	return findings
}

// checkUnclassifiedActions covers RMP-E007: every Allow'd IAM/STS action
// must be something the classifier recognizes.
func checkUnclassifiedActions(doc render.Document) []Finding {
	var findings []Finding

	seen := make(map[string]bool)

	for _, s := range doc.Statements {
		if s.Effect != "Allow" {
			continue
		}

		for _, action := range s.Action {
			if seen[action] || !isIAMOrSTS(action) {
				continue
			}

			seen[action] = true

			if classify.Classify(action) == classify.Unknown {
				findings = append(findings, Finding{
					Code:     "RMP-E007",
					Severity: Blocking,
					Message:  fmt.Sprintf("unclassified IAM/STS action %q is allowed in the deploy boundary", action),
				})
			}
		}
	}

	return findings
}

// checkTrustPolicyMFA covers RMP-E008.
func checkTrustPolicyMFA(doc render.TrustDocument) []Finding {
	var findings []Finding

	for _, s := range doc.Statement {
		if s.Effect != "Allow" || s.Action != "sts:AssumeRole" {
			continue
		}

		if s.Condition == nil || s.Condition.Bool["aws:MultiFactorAuthPresent"] != "true" {
			findings = append(findings, Finding{
				Code:     "RMP-E008",
				Severity: Blocking,
				Message:  "trust policy allows sts:AssumeRole without requiring aws:MultiFactorAuthPresent",
			})
		}
	}

	return findings
}

// checkAssumeRolePolicyMinimal covers RMP-E009: the source-principal
// policy must grant only sts:AssumeRole.
func checkAssumeRolePolicyMinimal(doc render.Document) []Finding {
	var findings []Finding

	for _, s := range doc.Statements {
		if s.Effect != "Allow" {
			continue
		}

		for _, action := range s.Action {
			if action != "sts:AssumeRole" {
				findings = append(findings, Finding{
					Code:     "RMP-E009",
					Severity: Blocking,
					Message:  fmt.Sprintf("source-principal policy grants %q, want only sts:AssumeRole", action),
				})
			}
		}
	}

	return findings
}

// checkWorkloadAllowlistWildcard covers the section 17.2 warning: a
// service wildcard (e.g. "s3:*") in the workload allowlist isn't blocked
// by config validation (only the specific IAM/STS/org/account wildcards
// are, per section 7.1) but is worth flagging.
func checkWorkloadAllowlistWildcard(cfg *config.Config) []Finding {
	var findings []Finding

	for _, action := range cfg.Policies.WorkloadBoundary.AllowedActions {
		if strings.HasSuffix(action, ":*") {
			findings = append(findings, Finding{
				Code:     "WORKLOAD_ALLOWLIST_WILDCARD",
				Severity: Warning,
				Message:  fmt.Sprintf("policies.workload_boundary.allowed_actions contains a service wildcard %q", action),
			})
		}
	}

	return findings
}

func hasDenyFor(doc render.Document, action string) bool {
	for _, s := range doc.Statements {
		if s.Effect == "Deny" && containsAction(s.Action, action) {
			return true
		}
	}

	return false
}

func containsAction(actions []string, action string) bool {
	for _, a := range actions {
		if a == action {
			return true
		}
	}

	return false
}

func isIAMOrSTS(action string) bool {
	return strings.HasPrefix(action, "iam:") || strings.HasPrefix(action, "sts:")
}

// conditionString and conditionStringList tolerate both a Go []string
// (the shape transform.DeployBoundary builds in memory) and the []any a
// fresh json.Unmarshal into map[string]any produces (the shape this
// package will see once artifacts are validated from disk), so this
// package works for either caller.
func conditionString(cond *render.Condition, key string) string {
	if cond == nil || cond.StringEquals == nil {
		return ""
	}

	s, _ := cond.StringEquals[key].(string)

	return s
}

func conditionStringList(cond *render.Condition, key string) []string {
	if cond == nil || cond.StringEquals == nil {
		return nil
	}

	switch v := cond.StringEquals[key].(type) {
	case []string:
		return v
	case []any:
		out := make([]string, 0, len(v))

		for _, item := range v {
			if s, ok := item.(string); ok {
				out = append(out, s)
			}
		}

		return out
	case string:
		return []string{v}
	default:
		return nil
	}
}

func roleARN(cfg *config.Config, path string) string {
	return fmt.Sprintf("arn:%s:iam::%s:role%s*", cfg.AWS.Partition, cfg.AWS.AccountID, path)
}

func namedPolicyARN(cfg *config.Config, path, name string) string {
	return fmt.Sprintf("arn:%s:iam::%s:policy%s%s", cfg.AWS.Partition, cfg.AWS.AccountID, path, name)
}
