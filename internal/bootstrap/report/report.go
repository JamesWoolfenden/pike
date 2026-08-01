// Package report builds report.json (spec section 19): the inputs,
// versions, findings, rewrites, warnings and artifact hashes for one
// generate run. It never calls time.Now() itself — GeneratedAt is a
// caller-supplied value — so Generate stays a pure function and its
// output is reproducible in tests.
package report

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/qj0r9j0vc2/rampart/internal/bootstrap/classify"
	"github.com/qj0r9j0vc2/rampart/internal/bootstrap/render"
	"github.com/qj0r9j0vc2/rampart/internal/bootstrap/validate"
)

const SchemaVersion = 1

// Report is report.json's top-level shape.
type Report struct {
	SchemaVersion int       `json:"schema_version"`
	GeneratedAt   string    `json:"generated_at"`
	Generator     Generator `json:"generator"`
	Scanner       Scanner   `json:"scanner"`
	Inputs        Inputs    `json:"inputs"`
	Summary       Summary   `json:"summary"`
	Rewrites      []Rewrite `json:"rewrites"`
	Warnings      []Warning `json:"warnings"`
	Artifacts     Artifacts `json:"artifacts"`
}

type Generator struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type Scanner struct {
	Name            string `json:"name"`
	MappingRevision string `json:"mapping_revision"`
}

type Inputs struct {
	TerraformDirectory string `json:"terraform_directory"`
	ConfigHash         string `json:"config_hash"`
}

type Summary struct {
	TotalActions      int `json:"total_actions"`
	IAMActions        int `json:"iam_actions"`
	EscalationActions int `json:"escalation_actions"`
	BlockingFindings  int `json:"blocking_findings"`
}

// Rewrite records one escalation-sensitive action that got a constrained
// replacement statement during deploy-boundary transformation, rather
// than being denied outright or passed through as a plain read.
type Rewrite struct {
	Action         string `json:"action"`
	Classification string `json:"classification"`
	Rule           string `json:"rule"`
}

type Warning struct {
	Code    string `json:"code"`
	Service string `json:"service,omitempty"`
}

type Artifacts struct {
	DeployPolicy     string `json:"deploy_policy"`
	DeployBoundary   string `json:"deploy_boundary"`
	WorkloadBoundary string `json:"workload_boundary"`
}

// rewriteRule maps the classify.Class values that get a constructive
// replacement statement (spec section 12) to that statement's Sid. Classes
// not listed here either pass through unchanged (Read) or are denied
// outright (BoundaryMutation, PolicyMutation, HumanPrincipal,
// Organization) — neither is a "rewrite" in report.json's sense.
var rewriteRule = map[classify.Class]string{
	classify.RoleCreate:        "CreateRoleWithRequiredBoundary",
	classify.RolePolicyWrite:   "ManageTerraformRoles",
	classify.PassRole:          "PassManagedRolesToApprovedServices",
	classify.ServiceLinkedRole: "CreateApprovedServiceLinkedRoles",
}

// Input bundles everything one generate run's report needs. GeneratedAt
// and GeneratorVersion are supplied by the caller rather than read from
// time.Now()/a package-level version var, so Generate has no hidden
// inputs and its output is exactly reproducible given the same Input.
type Input struct {
	GeneratedAt        time.Time
	GeneratorVersion   string
	ScannerName        string
	MappingRevision    string
	TerraformDirectory string
	ConfigBytes        []byte
	DeployPolicy       render.Document
	DeployBoundary     render.Document
	WorkloadBoundary   render.Document
	Findings           []validate.Finding
}

// Generate builds the report for one generate run.
func Generate(in Input) (Report, error) {
	deployPolicyHash, err := canonicalHash(in.DeployPolicy)
	if err != nil {
		return Report{}, fmt.Errorf("hashing deploy policy: %w", err)
	}

	deployBoundaryHash, err := canonicalHash(in.DeployBoundary)
	if err != nil {
		return Report{}, fmt.Errorf("hashing deploy boundary: %w", err)
	}

	workloadBoundaryHash, err := canonicalHash(in.WorkloadBoundary)
	if err != nil {
		return Report{}, fmt.Errorf("hashing workload boundary: %w", err)
	}

	blocking := 0

	for _, f := range in.Findings {
		if f.Severity == validate.Blocking {
			blocking++
		}
	}

	return Report{
		SchemaVersion: SchemaVersion,
		GeneratedAt:   in.GeneratedAt.UTC().Format(time.RFC3339),
		Generator:     Generator{Name: "rampart", Version: in.GeneratorVersion},
		Scanner:       Scanner{Name: in.ScannerName, MappingRevision: in.MappingRevision},
		Inputs: Inputs{
			TerraformDirectory: in.TerraformDirectory,
			ConfigHash:         "sha256:" + sha256Hex(in.ConfigBytes),
		},
		Summary:  summarize(in.DeployPolicy, blocking),
		Rewrites: rewrites(in.DeployPolicy),
		Warnings: warnings(in.Findings),
		Artifacts: Artifacts{
			DeployPolicy:     deployPolicyHash,
			DeployBoundary:   deployBoundaryHash,
			WorkloadBoundary: workloadBoundaryHash,
		},
	}, nil
}

func summarize(deployPolicy render.Document, blockingFindings int) Summary {
	totalActions := 0
	iamActions := 0

	for _, s := range deployPolicy.Statements {
		totalActions += len(s.Action)

		for _, action := range s.Action {
			if isIAMOrSTS(action) {
				iamActions++
			}
		}
	}

	escalation := escalationActions(deployPolicy)

	return Summary{
		TotalActions:      totalActions,
		IAMActions:        iamActions,
		EscalationActions: len(escalation),
		BlockingFindings:  blockingFindings,
	}
}

func rewrites(deployPolicy render.Document) []Rewrite {
	var out []Rewrite

	for _, action := range escalationActions(deployPolicy) {
		class := classify.Classify(action)

		rule, ok := rewriteRule[class]
		if !ok {
			continue
		}

		out = append(out, Rewrite{Action: action, Classification: string(class), Rule: rule})
	}

	return out
}

func warnings(findings []validate.Finding) []Warning {
	var out []Warning

	for _, f := range findings {
		if f.Severity == validate.Warning {
			out = append(out, Warning{Code: f.Code})
		}
	}

	return out
}

// escalationActions re-derives render.EscalationActions' result directly
// from a Document rather than a policy.Policy — the deploy policy has
// already been rendered by the time report.Generate runs, and pulling
// policy.Policy back in just to re-derive the same actions would add a
// dependency this package doesn't otherwise need.
func escalationActions(doc render.Document) []string {
	var found []string

	seen := make(map[string]bool)

	for _, s := range doc.Statements {
		for _, action := range s.Action {
			if seen[action] || !isIAMOrSTS(action) {
				continue
			}

			seen[action] = true

			if classify.Classify(action) != classify.Read {
				found = append(found, action)
			}
		}
	}

	return found
}

func isIAMOrSTS(action string) bool {
	return strings.HasPrefix(action, "iam:") || strings.HasPrefix(action, "sts:")
}

func canonicalHash(doc render.Document) (string, error) {
	data, err := json.Marshal(doc)
	if err != nil {
		return "", err
	}

	return "sha256:" + sha256Hex(data), nil
}

func sha256Hex(data []byte) string {
	sum := sha256.Sum256(data)

	return hex.EncodeToString(sum[:])
}
