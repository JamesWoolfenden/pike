package report_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/jameswoolfenden/pike/internal/bootstrap/render"
	"github.com/jameswoolfenden/pike/internal/bootstrap/report"
	"github.com/jameswoolfenden/pike/internal/bootstrap/validate"
)

func fixedTime() time.Time {
	return time.Date(2026, 7, 30, 0, 0, 0, 0, time.UTC)
}

func baseInput() report.Input {
	deployPolicy := render.Document{Version: "2012-10-17", Statements: []render.Statement{
		{Sid: "S0", Effect: "Allow", Action: []string{"ec2:DescribeInstances", "ec2:RunInstances"}, Resource: []string{"*"}},
		{Sid: "S1", Effect: "Allow", Action: []string{"iam:CreateRole", "iam:GetRole", "iam:PassRole"}, Resource: []string{"*"}},
	}}

	deployBoundary := render.Document{Version: "2012-10-17", Statements: []render.Statement{
		{Sid: "CreateRoleWithRequiredBoundary", Effect: "Allow", Action: []string{"iam:CreateRole"}, Resource: []string{"arn:aws:iam::123456789012:role/terraform-managed/*"}},
	}}

	workloadBoundary := render.Document{Version: "2012-10-17", Statements: []render.Statement{
		{Sid: "TerraformWorkloadBoundary", Effect: "Allow", Action: []string{"logs:PutLogEvents"}, Resource: []string{"*"}},
	}}

	return report.Input{
		GeneratedAt:        fixedTime(),
		GeneratorVersion:   "0.1.0",
		ScannerName:        "pike",
		MappingRevision:    "rev1",
		TerraformDirectory: "./infra",
		ConfigBytes:        []byte("version: 1\n"),
		DeployPolicy:       deployPolicy,
		DeployBoundary:     deployBoundary,
		WorkloadBoundary:   workloadBoundary,
		Findings: []validate.Finding{
			{Code: "WORKLOAD_ALLOWLIST_WILDCARD", Severity: validate.Warning, Message: "warn"},
		},
	}
}

func TestGenerate(t *testing.T) {
	t.Parallel()

	r, err := report.Generate(baseInput())
	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}

	if r.SchemaVersion != report.SchemaVersion {
		t.Errorf("SchemaVersion = %d, want %d", r.SchemaVersion, report.SchemaVersion)
	}

	if r.GeneratedAt != "2026-07-30T00:00:00Z" {
		t.Errorf("GeneratedAt = %q, want 2026-07-30T00:00:00Z", r.GeneratedAt)
	}

	if r.Generator.Name != "pike" || r.Generator.Version != "0.1.0" {
		t.Errorf("Generator = %+v, want {pike 0.1.0}", r.Generator)
	}

	if r.Scanner.Name != "pike" || r.Scanner.MappingRevision != "rev1" {
		t.Errorf("Scanner = %+v, want {pike rev1}", r.Scanner)
	}

	if r.Inputs.TerraformDirectory != "./infra" {
		t.Errorf("Inputs.TerraformDirectory = %q, want ./infra", r.Inputs.TerraformDirectory)
	}

	if r.Inputs.ConfigHash == "" || r.Inputs.ConfigHash[:7] != "sha256:" {
		t.Errorf("Inputs.ConfigHash = %q, want a sha256: prefix", r.Inputs.ConfigHash)
	}

	// total_actions: 2 (ec2) + 3 (iam) = 5. iam_actions: CreateRole, GetRole, PassRole = 3.
	// escalation_actions: CreateRole, PassRole (GetRole is a read) = 2.
	if r.Summary.TotalActions != 5 {
		t.Errorf("Summary.TotalActions = %d, want 5", r.Summary.TotalActions)
	}

	if r.Summary.IAMActions != 3 {
		t.Errorf("Summary.IAMActions = %d, want 3", r.Summary.IAMActions)
	}

	if r.Summary.EscalationActions != 2 {
		t.Errorf("Summary.EscalationActions = %d, want 2", r.Summary.EscalationActions)
	}

	if r.Summary.BlockingFindings != 0 {
		t.Errorf("Summary.BlockingFindings = %d, want 0 (only a warning finding was given)", r.Summary.BlockingFindings)
	}

	wantRewrites := map[string]string{
		"iam:CreateRole": "CreateRoleWithRequiredBoundary",
		"iam:PassRole":   "PassManagedRolesToApprovedServices",
	}

	if len(r.Rewrites) != len(wantRewrites) {
		t.Fatalf("Rewrites = %+v, want entries for %v", r.Rewrites, wantRewrites)
	}

	for _, rw := range r.Rewrites {
		if wantRule, ok := wantRewrites[rw.Action]; !ok || wantRule != rw.Rule {
			t.Errorf("Rewrite for %q = %+v, want rule %q", rw.Action, rw, wantRewrites[rw.Action])
		}
	}

	if len(r.Warnings) != 1 || r.Warnings[0].Code != "WORKLOAD_ALLOWLIST_WILDCARD" {
		t.Errorf("Warnings = %+v, want one WORKLOAD_ALLOWLIST_WILDCARD entry", r.Warnings)
	}

	for name, hash := range map[string]string{
		"DeployPolicy":     r.Artifacts.DeployPolicy,
		"DeployBoundary":   r.Artifacts.DeployBoundary,
		"WorkloadBoundary": r.Artifacts.WorkloadBoundary,
	} {
		if hash == "" || hash[:7] != "sha256:" {
			t.Errorf("Artifacts.%s = %q, want a sha256: prefix", name, hash)
		}
	}
}

func TestGenerate_BlockingFindingsCounted(t *testing.T) {
	t.Parallel()

	in := baseInput()
	in.Findings = []validate.Finding{
		{Code: "RMP-E001", Severity: validate.Blocking},
		{Code: "RMP-E002", Severity: validate.Blocking},
		{Code: "WORKLOAD_ALLOWLIST_WILDCARD", Severity: validate.Warning},
	}

	r, err := report.Generate(in)
	if err != nil {
		t.Fatal(err)
	}

	if r.Summary.BlockingFindings != 2 {
		t.Errorf("Summary.BlockingFindings = %d, want 2", r.Summary.BlockingFindings)
	}

	if len(r.Warnings) != 1 {
		t.Errorf("Warnings = %+v, want exactly the one warning-severity finding", r.Warnings)
	}
}

func TestGenerate_Deterministic(t *testing.T) {
	t.Parallel()

	in := baseInput()

	first, err := report.Generate(in)
	if err != nil {
		t.Fatal(err)
	}

	firstJSON, err := json.Marshal(first)
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		r, err := report.Generate(in)
		if err != nil {
			t.Fatal(err)
		}

		got, err := json.Marshal(r)
		if err != nil {
			t.Fatal(err)
		}

		if string(got) != string(firstJSON) {
			t.Fatalf("run %d: Generate() rendered differently:\n%s\nvs\n%s", i, got, firstJSON)
		}
	}
}
