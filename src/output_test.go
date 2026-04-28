package pike

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestWriteOutput_InvalidFormat covers the branch where outFile is empty and
// outputType is not terraform or json — returns tfPolicyFormatError.
func TestWriteOutput_InvalidFormat(t *testing.T) {
	t.Parallel()

	policy := OutputPolicy{AWS: AwsOutput{JSONOut: `{"Version":"2012-10-17"}`}}
	err := WriteOutput(policy, "yaml", t.TempDir(), "")
	if err == nil {
		t.Fatal("WriteOutput() with unknown format expected error, got nil")
	}
}

// TestWriteOutput_RoleFile covers the branch that writes aws_iam_role.terraform_pike.tf
// alongside the policy file when AWS.Terraform is non-empty.
func TestWriteOutput_RoleFile(t *testing.T) {
	t.Parallel()

	policy := OutputPolicy{
		AWS: AwsOutput{
			Terraform: `resource "aws_iam_policy" "pike" { name = "pike" }`,
			JSONOut:   `{"Version":"2012-10-17"}`,
		},
	}

	tmp := t.TempDir()
	if err := WriteOutput(policy, "terraform", tmp, ""); err != nil {
		t.Fatalf("WriteOutput() error = %v", err)
	}

	roleFile := filepath.Join(tmp, ".pike", "aws_iam_role.terraform_pike.tf")
	data, err := os.ReadFile(roleFile)
	if err != nil {
		t.Fatalf("expected role file %s to exist: %v", roleFile, err)
	}
	if len(data) == 0 {
		t.Error("role file is empty")
	}
}

// TestWriteOutput_ExplicitOutfile covers the outFile != "" fast path.
func TestWriteOutput_ExplicitOutfile(t *testing.T) {
	t.Parallel()

	policy := OutputPolicy{AWS: AwsOutput{JSONOut: `{"Version":"2012-10-17"}`}}
	outFile := filepath.Join(t.TempDir(), "my_policy.json")

	if err := WriteOutput(policy, "json", "", outFile); err != nil {
		t.Fatalf("WriteOutput() error = %v", err)
	}

	data, err := os.ReadFile(outFile)
	if err != nil {
		t.Fatalf("expected output file: %v", err)
	}
	if !strings.Contains(string(data), "2012-10-17") {
		t.Errorf("output file content %q missing expected policy version", string(data))
	}
}
