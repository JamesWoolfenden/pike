package scan_test

import (
	"context"
	"testing"

	"github.com/jameswoolfenden/pike/internal/scan"
)

func TestPikeAnalyzer_Scan(t *testing.T) {
	t.Parallel()

	var analyzer scan.PikeAnalyzer

	p, err := analyzer.Scan(context.Background(), "testdata/aws_security_group")
	if err != nil {
		t.Fatalf("Scan() error = %v, want nil", err)
	}

	if p.Empty() {
		t.Fatal("Scan() returned an empty policy, want the aws_security_group permissions")
	}

	actions := p.Actions()

	want := "ec2:CreateSecurityGroup"
	found := false

	for _, a := range actions {
		if a == want {
			found = true

			break
		}
	}

	if !found {
		t.Errorf("Scan() actions = %v, want it to contain %q", actions, want)
	}

	for _, s := range p.Statements {
		if len(s.Resources) != 1 || s.Resources[0] != "*" {
			t.Errorf("Statement.Resources = %v, want [\"*\"] (no AWS account context is available at scan time)", s.Resources)
		}
	}
}

func TestPikeAnalyzer_Scan_EmptyDirectory(t *testing.T) {
	t.Parallel()

	var analyzer scan.PikeAnalyzer

	p, err := analyzer.Scan(context.Background(), t.TempDir())
	if err != nil {
		t.Fatalf("Scan() error = %v, want nil for a directory with no Terraform files", err)
	}

	if !p.Empty() {
		t.Errorf("Scan() = %+v, want an empty policy", p)
	}
}

func TestPikeAnalyzer_Scan_ContextCanceled(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	var analyzer scan.PikeAnalyzer

	if _, err := analyzer.Scan(ctx, "testdata/aws_security_group"); err == nil {
		t.Fatal("Scan() error = nil, want context.Canceled to be surfaced")
	}
}
