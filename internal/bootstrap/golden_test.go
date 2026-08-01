// Package bootstrap_test holds end-to-end golden-fixture tests that
// exercise the full generate pipeline (scan -> render deploy-policy ->
// transform deploy-boundary -> render workload-boundary/trust-policy/
// assume-role-policy -> validate -> report) against realistic input, spec
// section 21.3.
package bootstrap_test

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/jameswoolfenden/pike/internal/bootstrap/config"
	"github.com/jameswoolfenden/pike/internal/bootstrap/render"
	"github.com/jameswoolfenden/pike/internal/bootstrap/report"
	"github.com/jameswoolfenden/pike/internal/bootstrap/transform"
	"github.com/jameswoolfenden/pike/internal/bootstrap/validate"
	"github.com/jameswoolfenden/pike/internal/policy"
	"github.com/jameswoolfenden/pike/internal/scan"
)

// goldenTime is the fixed GeneratedAt every fixture's report.json was
// captured with, so report generation stays comparable across runs.
var goldenTime = time.Date(2026, 7, 30, 0, 0, 0, 0, time.UTC)

// syntheticPolicies holds the scan input for fixtures with no terraform/
// directory — see each fixture's README.md for why a hand-built Policy
// was used instead of a scanned one.
var syntheticPolicies = map[string]policy.Policy{
	"unknown-iam-action": {Statements: []policy.Statement{{Actions: []string{"iam:UpdateAssumeRolePolicy"}, Resources: []string{"*"}}}},
	"managed-policy":     {Statements: []policy.Statement{{Actions: []string{"iam:CreatePolicyVersion"}, Resources: []string{"*"}}}},
}

func TestGolden(t *testing.T) {
	fixtures := []string{
		"no-iam",
		"ecs-task-role",
		"service-linked-role",
		"managed-policy",
		"unknown-iam-action",
	}

	for _, name := range fixtures {
		name := name
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			runGoldenFixture(t, name)
		})
	}
}

func runGoldenFixture(t *testing.T, name string) {
	t.Helper()

	dir := filepath.Join("..", "..", "testdata", "bootstrap", name)

	configBytes, err := os.ReadFile(filepath.Join(dir, "pike.yaml"))
	if err != nil {
		t.Fatalf("reading pike.yaml: %v", err)
	}

	cfg, err := config.Load(filepath.Join(dir, "pike.yaml"))
	if err != nil {
		t.Fatalf("config.Load() error = %v", err)
	}

	p, ok := syntheticPolicies[name]
	if !ok {
		var a scan.PikeAnalyzer

		p, err = a.Scan(context.Background(), filepath.Join(dir, "terraform"))
		if err != nil {
			t.Fatalf("Scan() error = %v", err)
		}
	}

	deployPolicy := render.DeployPolicy(p)
	assertGoldenJSON(t, filepath.Join(dir, "expected", "deploy-policy.json"), deployPolicy)

	deployBoundary, err := transform.DeployBoundary(p, cfg)
	if name == "unknown-iam-action" {
		var target *transform.UnknownIAMActionError
		if !errors.As(err, &target) {
			t.Fatalf("DeployBoundary() error = %v, want *UnknownIAMActionError (this fixture only has an expected deploy-policy.json — see its README.md)", err)
		}

		return
	}

	if err != nil {
		t.Fatalf("DeployBoundary() error = %v", err)
	}

	assertGoldenJSON(t, filepath.Join(dir, "expected", "deploy-boundary.json"), deployBoundary)

	workloadBoundary, err := render.WorkloadBoundary(cfg)
	if err != nil {
		t.Fatalf("WorkloadBoundary() error = %v", err)
	}

	assertGoldenJSON(t, filepath.Join(dir, "expected", "workload-boundary.json"), workloadBoundary)

	trustPolicy, err := render.TrustPolicy(cfg)
	if err != nil {
		t.Fatalf("TrustPolicy() error = %v", err)
	}

	assertGoldenJSON(t, filepath.Join(dir, "expected", "trust-policy.json"), trustPolicy)

	assumeRolePolicy := render.AssumeRolePolicy(cfg)
	assertGoldenJSON(t, filepath.Join(dir, "expected", "assume-role-policy.json"), assumeRolePolicy)

	findings := validate.Validate(validate.Artifacts{
		DeployBoundary:   deployBoundary,
		TrustPolicy:      trustPolicy,
		AssumeRolePolicy: assumeRolePolicy,
	}, cfg)

	for _, f := range findings {
		if f.Severity == validate.Blocking {
			t.Errorf("unexpected blocking finding on a golden fixture: %+v", f)
		}
	}

	r, err := report.Generate(report.Input{
		GeneratedAt:        goldenTime,
		GeneratorVersion:   "0.1.0",
		ScannerName:        "pike",
		MappingRevision:    "test",
		TerraformDirectory: "./terraform",
		ConfigBytes:        configBytes,
		DeployPolicy:       deployPolicy,
		DeployBoundary:     deployBoundary,
		WorkloadBoundary:   workloadBoundary,
		Findings:           findings,
	})
	if err != nil {
		t.Fatalf("report.Generate() error = %v", err)
	}

	assertGoldenJSON(t, filepath.Join(dir, "expected", "report.json"), r)
}

// assertGoldenJSON compares got against the JSON file at path, both
// normalized through json.Unmarshal into `any` first so formatting
// differences (whitespace, key order in the source file) don't cause
// false failures — only structural differences do.
func assertGoldenJSON(t *testing.T, path string, got any) {
	t.Helper()

	want, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("reading golden file %s: %v", path, err)
	}

	gotBytes, err := json.Marshal(got)
	if err != nil {
		t.Fatalf("marshaling actual output for %s: %v", path, err)
	}

	var wantNormalized, gotNormalized any

	if err := json.Unmarshal(want, &wantNormalized); err != nil {
		t.Fatalf("parsing golden file %s: %v", path, err)
	}

	if err := json.Unmarshal(gotBytes, &gotNormalized); err != nil {
		t.Fatalf("parsing actual output for %s: %v", path, err)
	}

	wantCanonical, _ := json.Marshal(wantNormalized)
	gotCanonical, _ := json.Marshal(gotNormalized)

	if string(gotCanonical) != string(wantCanonical) {
		gotPretty, _ := json.MarshalIndent(gotNormalized, "", "  ")
		t.Errorf("%s mismatch.\ngot:\n%s\n\nwant (golden file):\n%s", path, gotPretty, want)
	}
}
