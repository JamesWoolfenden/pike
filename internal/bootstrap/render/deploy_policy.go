// Package render converts Pike's internal Policy model into the JSON
// artifacts Pike writes to disk (spec section 8), starting with
// deploy-policy.json (section 11).
package render

import (
	"sort"
	"strconv"
	"strings"

	"github.com/jameswoolfenden/pike/internal/bootstrap/classify"
	"github.com/jameswoolfenden/pike/internal/policy"
)

// policyVersion is the fixed AWS IAM policy language version.
const policyVersion = "2012-10-17"

// Document is a rendered AWS IAM policy document.
type Document struct {
	Version    string      `json:"Version"`
	Statements []Statement `json:"Statement"`
}

// Statement is a single Allow/Deny block within a Document.
type Statement struct {
	Sid       string     `json:"Sid"`
	Effect    string     `json:"Effect"`
	Action    []string   `json:"Action"`
	Resource  []string   `json:"Resource"`
	Condition *Condition `json:"Condition,omitempty"`
}

// Condition is an IAM policy condition block. Only StringEquals is
// needed for the conditions Pike generates today: CreateRole's
// PermissionsBoundary check and the PassRole/CreateServiceLinkedRole
// service allowlists (spec section 12). Values are either a single
// string or a string list; encoding/json sorts map keys automatically,
// so this stays canonical (section 8.1) without extra work.
type Condition struct {
	StringEquals map[string]any `json:"StringEquals,omitempty"`
}

// DeployPolicy renders deploy-policy.json: the permissions Terraform
// requests, preserved as scanned (spec section 11). It does not apply any
// security restriction — that's deploy-boundary.json's job, kept in a
// separate stage so the two artifacts stay independently reviewable
// ("Deploy policy = what Terraform requests, Deploy boundary = maximum
// safe delegation").
//
// Output is canonical (spec section 8.1): actions and resources within
// each statement are sorted, and statements are ordered by their first
// action, so identical scan input always renders identical bytes
// regardless of any nondeterminism upstream in how Policy was built.
func DeployPolicy(p policy.Policy) Document {
	statements := make([]Statement, 0, len(p.Statements))

	for i, s := range p.Statements {
		actions := sortedUnique(s.Actions)
		if len(actions) == 0 {
			continue
		}

		statements = append(statements, Statement{
			Sid:      "DeployPolicy" + strconv.Itoa(i),
			Effect:   "Allow",
			Action:   actions,
			Resource: sortedUnique(s.Resources),
		})
	}

	sort.Slice(statements, func(i, j int) bool {
		return firstAction(statements[i]) < firstAction(statements[j])
	})

	for i := range statements {
		statements[i].Sid = "DeployPolicy" + strconv.Itoa(i)
	}

	return Document{Version: policyVersion, Statements: statements}
}

// EscalationActions returns the deduplicated, sorted set of IAM/STS
// actions in p that are not plain reads — every action spec section 3
// calls an "escalation action" (one that can grant, pass, mutate or
// indirectly obtain additional privileges). Section 11 requires these be
// recorded in report.json; this is what that stage will read from.
func EscalationActions(p policy.Policy) []string {
	seen := make(map[string]bool)

	var found []string

	for _, action := range p.Actions() {
		if !isIAMOrSTS(action) {
			continue
		}

		if classify.Classify(action) == classify.Read {
			continue
		}

		if seen[action] {
			continue
		}

		seen[action] = true

		found = append(found, action)
	}

	sort.Strings(found)

	return found
}

func isIAMOrSTS(action string) bool {
	return strings.HasPrefix(action, "iam:") || strings.HasPrefix(action, "sts:")
}

// sortedUnique returns the sorted, deduplicated contents of s. Pike's scan
// output can list the same action more than once within a statement (e.g.
// several attributes on one resource all requiring iam:GetRole); a
// rendered policy document shouldn't repeat it.
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

func firstAction(s Statement) string {
	if len(s.Action) == 0 {
		return ""
	}

	return s.Action[0]
}
