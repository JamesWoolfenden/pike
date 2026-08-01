// Package policy defines Rampart's provider-neutral internal representation
// of the permissions a piece of infrastructure-as-code requires. It exists
// so the bootstrap-generation packages (classify, transform, validate, ...)
// depend on this stable shape rather than reaching into Pike's own
// scan-result types directly — keeping Pike's mappings separable for
// upstream synchronization (see spec section 20.2).
package policy

// Policy is the permissions required to deploy a Terraform/OpenTofu
// configuration, as produced by an Analyzer (see internal/scan).
type Policy struct {
	Statements []Statement
}

// Statement groups one or more actions under a shared resource scope. A
// Policy returned from scanning source code is always an implicit "Allow" —
// deny semantics are introduced later, during deploy-boundary
// transformation, not at scan time.
type Statement struct {
	Actions   []string
	Resources []string
}

// Actions returns every action across all statements, deduplicated but
// unsorted (callers that need a stable order should sort the result).
func (p Policy) Actions() []string {
	seen := make(map[string]bool)

	var actions []string

	for _, statement := range p.Statements {
		for _, action := range statement.Actions {
			if seen[action] {
				continue
			}

			seen[action] = true

			actions = append(actions, action)
		}
	}

	return actions
}

// Empty reports whether the policy has no statements, i.e. the scanned
// source required no permissions.
func (p Policy) Empty() bool {
	return len(p.Statements) == 0
}
