// Package bootstrap wires the generate -> validate -> report pipeline
// (config, classify, render, transform, validate, report) into a single
// entry point the CLI's "bootstrap" command drives. This is the AWS-only,
// EXPERIMENTAL IAM deploy-role bootstrap: it prevents a Terraform deploy
// role's iam:CreateRole/PassRole access from being used to escalate to
// account admin, but it is not a least-privilege policy generator - every
// non-IAM action pike detects is still scoped to Resource ["*"], since
// generation is static analysis only and never inspects live AWS state.
package bootstrap

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/jameswoolfenden/pike/internal/bootstrap/config"
	"github.com/jameswoolfenden/pike/internal/bootstrap/render"
	"github.com/jameswoolfenden/pike/internal/bootstrap/report"
	"github.com/jameswoolfenden/pike/internal/bootstrap/transform"
	"github.com/jameswoolfenden/pike/internal/bootstrap/validate"
	"github.com/jameswoolfenden/pike/internal/scan"
	pike "github.com/jameswoolfenden/pike/src"
)

// defaultOutputSubdir is where generated artifacts land under the scanned
// directory when pike.yaml's output.directory is empty, alongside pike's
// existing .pike output convention (pike.generated_policy.tf etc).
const defaultOutputSubdir = ".pike/bootstrap"

// artifactFilenames are the six files a generate run writes.
var artifactFilenames = []string{
	"deploy-policy.json",
	"deploy-boundary.json",
	"workload-boundary.json",
	"trust-policy.json",
	"assume-role-policy.json",
	"report.json",
}

// Run scans directory for AWS IAM permissions, loads configPath (pike.yaml
// by default), and generates the deploy-policy/deploy-boundary/workload-
// boundary/trust-policy/assume-role-policy/report artifacts. Artifacts are
// always written, even when validation produces blocking findings, so the
// caller can inspect what's wrong - but Run returns a non-nil error in
// that case (and CI callers should treat that as generation having
// failed).
func Run(directory, configPath string) error {
	configBytes, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("reading %s: %w", configPath, err)
	}

	cfg, err := config.Parse(configBytes)
	if err != nil {
		return fmt.Errorf("%s: %w", configPath, err)
	}

	var analyzer scan.PikeAnalyzer

	p, err := analyzer.Scan(context.Background(), directory)
	if err != nil {
		return fmt.Errorf("scanning %s: %w", directory, err)
	}

	deployPolicy := render.DeployPolicy(p)

	deployBoundary, err := transform.DeployBoundary(p, cfg)
	if err != nil {
		return fmt.Errorf("generating deploy boundary: %w", err)
	}

	workloadBoundary, err := render.WorkloadBoundary(cfg)
	if err != nil {
		return fmt.Errorf("generating workload boundary: %w", err)
	}

	trustPolicy, err := render.TrustPolicy(cfg)
	if err != nil {
		return fmt.Errorf("generating trust policy: %w", err)
	}

	assumeRolePolicy := render.AssumeRolePolicy(cfg)

	findings := validate.Validate(validate.Artifacts{
		DeployBoundary:   deployBoundary,
		TrustPolicy:      trustPolicy,
		AssumeRolePolicy: assumeRolePolicy,
	}, cfg)

	r, err := report.Generate(report.Input{
		GeneratedAt:        time.Now().UTC(),
		GeneratorVersion:   pike.Version,
		ScannerName:        "pike",
		MappingRevision:    pike.Version,
		TerraformDirectory: directory,
		ConfigBytes:        configBytes,
		DeployPolicy:       deployPolicy,
		DeployBoundary:     deployBoundary,
		WorkloadBoundary:   workloadBoundary,
		Findings:           findings,
	})
	if err != nil {
		return fmt.Errorf("generating report: %w", err)
	}

	outputDir, err := resolveOutputDir(directory, cfg)
	if err != nil {
		return err
	}

	if !cfg.Output.Overwrite {
		if err := refuseIfArtifactsExist(outputDir); err != nil {
			return err
		}
	}

	if err := os.MkdirAll(outputDir, 0o750); err != nil {
		return fmt.Errorf("creating %s: %w", outputDir, err)
	}

	artifacts := map[string]any{
		"deploy-policy.json":      deployPolicy,
		"deploy-boundary.json":    deployBoundary,
		"workload-boundary.json":  workloadBoundary,
		"trust-policy.json":       trustPolicy,
		"assume-role-policy.json": assumeRolePolicy,
		"report.json":             r,
	}

	for _, name := range artifactFilenames {
		if err := writeJSON(filepath.Join(outputDir, name), artifacts[name]); err != nil {
			return err
		}
	}

	fmt.Printf("pike bootstrap: wrote %d artifacts to %s\n", len(artifactFilenames), outputDir)

	var blocking int

	for _, f := range findings {
		if f.Severity == validate.Blocking {
			blocking++

			fmt.Printf("  BLOCKING [%s] %s\n", f.Code, f.Message)
		}
	}

	if blocking > 0 {
		return fmt.Errorf("%d blocking validation finding(s), see %s", blocking, filepath.Join(outputDir, "report.json"))
	}

	return nil
}

// resolveOutputDir honors pike.yaml's output.directory when set (relative
// paths are resolved against the scanned directory, matching how
// directory-scoped output already works elsewhere in pike), defaulting to
// defaultOutputSubdir otherwise.
func resolveOutputDir(directory string, cfg *config.Config) (string, error) {
	if cfg.Output.Directory == "" {
		return filepath.Join(directory, defaultOutputSubdir), nil
	}

	if filepath.IsAbs(cfg.Output.Directory) {
		return cfg.Output.Directory, nil
	}

	return filepath.Join(directory, cfg.Output.Directory), nil
}

// refuseIfArtifactsExist protects existing artifacts from a silent
// overwrite when output.overwrite is false (the config default) - the
// caller must opt in before a re-run replaces prior output.
func refuseIfArtifactsExist(outputDir string) error {
	for _, name := range artifactFilenames {
		path := filepath.Join(outputDir, name)
		if _, err := os.Stat(path); err == nil {
			return fmt.Errorf("%s already exists and output.overwrite is false", path)
		}
	}

	return nil
}

func writeJSON(path string, v any) error {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf("marshaling %s: %w", path, err)
	}

	b = append(b, '\n')

	if err := os.WriteFile(path, b, 0o600); err != nil {
		return fmt.Errorf("writing %s: %w", path, err)
	}

	return nil
}
