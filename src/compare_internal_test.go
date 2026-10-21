package pike

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	diff "github.com/yudai/gojsondiff"
	gcpiam "google.golang.org/api/iam/v1"
)

type mockDiff struct {
	diff.Diff
}

func (m mockDiff) Modified() bool {
	return true
}

func TestCompareIAMPolicy(t *testing.T) {
	t.Parallel()

	type args struct {
		Policy    string
		OldPolicy string
	}

	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			"same",
			args{
				"{\"Statement\":[{\"Action\":[\"cognito-idp:ListUserPoolClients\",\"cognito-idp:GetSigningCertificate\",\"cognito-idp:DescribeUserPoolClient\"],\"Effect\":\"Allow\",\"Resource\":\"*\",\"Sid\":\"0\"}],\"Version\":\"2012-10-17\"}",
				"{\"Statement\":[{\"Action\":[\"cognito-idp:ListUserPoolClients\",\"cognito-idp:GetSigningCertificate\",\"cognito-idp:DescribeUserPoolClient\"],\"Effect\":\"Allow\",\"Resource\":\"*\",\"Sid\":\"0\"}],\"Version\":\"2012-10-17\"}",
			},
			true,
			false,
		},
		{
			"different",
			args{
				"{\"Statement\":[{\"Action\":[\"cognito-idp:ListUserPoolClients\",\"cognito-idp:GetSigningCertificate\",\"cognito-idp:DescribeUserPoolClient\"],\"Effect\":\"Allow\",\"Resource\":\"*\",\"Sid\":\"0\"}],\"Version\":\"2012-10-17\"}",
				"{\"Statement\":[{\"Action\":[\"cognito-idp:ListUserPoolClients\",\"cognito-idp:GetSigningCertificate\",\"cognito-idp:bogus\",\"cognito-idp:DescribeUserPoolClient\"],\"Effect\":\"Allow\",\"Resource\":\"*\",\"Sid\":\"0\"}],\"Version\":\"2012-10-17\"}",
			},
			false,
			false,
		},
		{
			"not-json",
			args{
				"guff",
				"{\"Statement\":[{\"Action\":[\"cognito-idp:ListUserPoolClients\",\"cognito-idp:GetSigningCertificate\",\"cognito-idp:bogus\",\"cognito-idp:DescribeUserPoolClient\"],\"Effect\":\"Allow\",\"Resource\":\"*\",\"Sid\":\"0\"}],\"Version\":\"2012-10-17\"}",
			},
			false,
			true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := compareIAMPolicy(tt.args.Policy, tt.args.OldPolicy)
			if (err != nil) != tt.wantErr {
				t.Errorf("CompareIAMPolicy() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if got != tt.want {
				t.Errorf("CompareIAMPolicy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShowDifferences(t *testing.T) {
	tests := []struct {
		name        string
		policy      string
		diff        diff.Diff
		wantBool    bool
		wantErr     bool
		description string
	}{
		{
			name:        "Invalid JSON policy",
			policy:      `{invalid-json}`,
			diff:        &mockDiff{},
			wantBool:    false,
			wantErr:     true,
			description: "Should return error for invalid JSON",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBool, err := showDifferences(tt.policy, tt.diff)

			if (err != nil) != tt.wantErr {
				t.Errorf("ShowDifferences() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotBool != tt.wantBool {
				t.Errorf("ShowDifferences() = %v, want %v", gotBool, tt.wantBool)
			}
		})
	}
}

func TestInputValidationCompare(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "pike-test-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	tests := []struct {
		name      string
		directory string
		arn       string
		wantBool  bool
		wantErr   error
	}{
		{
			name:      "empty directory",
			directory: "",
			arn:       "arn:aws:iam::123456789012:policy/test",
			wantBool:  false,
			wantErr:   &emptyDirectoryError{},
		},
		{
			name:      "non-existent directory",
			directory: filepath.Join(tmpDir, "nonexistent"),
			arn:       "arn:aws:iam::123456789012:policy/test",
			wantBool:  false,
			wantErr:   &directoryNotFoundError{filepath.Join(tmpDir, "nonexistent")},
		},
		{
			name:      "empty ARN",
			directory: tmpDir,
			arn:       "",
			wantBool:  false,
			wantErr:   &arnEmptyError{},
		},
		{
			name:      "invalid ARN format",
			directory: tmpDir,
			arn:       "invalid:arn",
			wantBool:  false,
			wantErr:   &invalidARNError{"invalid:arn"},
		},
		{
			name:      "valid inputs",
			directory: tmpDir,
			arn:       "arn:aws:iam::123456789012:policy/test",
			wantBool:  true,
			wantErr:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBool, gotErr := inputValidationCompare(tt.directory, tt.arn)
			if gotBool != tt.wantBool {
				t.Errorf("inputValidationCompare() bool = %v, want %v", gotBool, tt.wantBool)
			}
			if (gotErr == nil && tt.wantErr != nil) || (gotErr != nil && tt.wantErr == nil) {
				t.Errorf("inputValidationCompare() error = %v, want %v", gotErr, tt.wantErr)
			}
			if gotErr != nil && tt.wantErr != nil && gotErr.Error() != tt.wantErr.Error() {
				t.Errorf("inputValidationCompare() error = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestGetCloudFromRole(t *testing.T) {
	tests := []struct {
		name     string
		arn      string
		expected string
	}{
		{
			name:     "AWS ARN",
			arn:      "arn:aws:iam::123456789012:role/MyRole",
			expected: "aws",
		},
		{
			name:     "GCP project role",
			arn:      "projects/my-project/roles/my-role",
			expected: "gcp",
		},
		{
			name:     "Unknown format",
			arn:      "invalid-role-format",
			expected: "unknown",
		},
		{
			name:     "Empty string",
			arn:      "",
			expected: "unknown",
		},
		{
			name:     "Partial AWS ARN",
			arn:      "arn:something",
			expected: "aws",
		},
		{
			name:     "Partial GCP format",
			arn:      "projects/test",
			expected: "gcp",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getCloudFromRole(tt.arn)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestVerifyGCPRole(t *testing.T) {
	tests := []struct {
		name        string
		role        string
		expectError bool
	}{
		{
			name:        "valid GCP role format",
			role:        "projects/my-project/roles/my-custom-role",
			expectError: false,
		},
		{
			name:        "valid GCP role with numbers and hyphens",
			role:        "projects/my-project-123/roles/custom-role-456",
			expectError: false,
		},
		{
			name:        "valid GCP role with underscores",
			role:        "projects/my_project/roles/my_role",
			expectError: false,
		},
		{
			name:        "invalid role - missing projects prefix",
			role:        "my-project/roles/my-role",
			expectError: true,
		},
		{
			name:        "invalid role - missing roles section",
			role:        "projects/my-project/my-role",
			expectError: true,
		},
		{
			name:        "invalid role - empty project name",
			role:        "projects//roles/my-role",
			expectError: true,
		},
		{
			name:        "invalid role - empty role name",
			role:        "projects/my-project/roles/",
			expectError: true,
		},
		{
			name:        "invalid role - completely empty",
			role:        "",
			expectError: true,
		},
		{
			name:        "invalid role - wrong format",
			role:        "organizations/123456789/roles/my-role",
			expectError: true,
		},
		{
			name:        "invalid role - predefined role format",
			role:        "roles/viewer",
			expectError: true,
		},
		{
			name:        "invalid role - trailing whitespace in project",
			role:        "projects/my-project /roles/my-role",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := verifyGCPRole(tt.role)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error for role %q, but got nil", tt.role)
				} else {
					// Verify it's the correct error type
					if _, ok := err.(*gcpRoleNotVerified); !ok {
						t.Errorf("expected gcpRoleNotVerified error type, got %T", err)
					}
				}
			} else {
				if err != nil {
					t.Errorf("expected no error for role %q, but got: %v", tt.role, err)
				}
			}
		})
	}
}

func TestVerifyGCPRole_ErrorMessage(t *testing.T) {
	invalidRole := "invalid-role-format"
	err := verifyGCPRole(invalidRole)

	if err == nil {
		t.Fatal("expected error but got nil")
	}

	gcpErr, ok := err.(*gcpRoleNotVerified)
	if !ok {
		t.Fatalf("expected gcpRoleNotVerified error type, got %T", err)
	}

	if gcpErr.role != invalidRole {
		t.Errorf("expected error to contain role %q, got %q", invalidRole, gcpErr.role)
	}
}

func Test_gcpIAMRoleError_Error(t *testing.T) {
	type fields struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"invoke", fields{err: &gcpRoleNotVerified{role: "test"}}, `IAM Role Error: GCP role "test" not verified: expected projects/{PROJECT}/roles/{ROLE}, organizations/{ORG}/roles/{ROLE}, or roles/{ROLE}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &gcpIAMRoleError{
				err: tt.fields.err,
			}
			assert.Equalf(t, tt.want, m.Error(), "Error()")
		})
	}
}

func Test_gcpRoleNotFound_Error(t *testing.T) {
	type fields struct {
		role string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Invoke", fields{role: "test"}, "IAM Role Error: test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &gcpRoleNotFound{
				role: tt.fields.role,
			}
			assert.Equalf(t, tt.want, e.Error(), "Error()")
		})
	}
}

func Test_compareDifferenceError_Error(t *testing.T) {
	type fields struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Invoke", fields{errors.New("test")}, "compare difference failed: test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &compareDifferenceError{
				err: tt.fields.err,
			}
			assert.Equalf(t, tt.want, m.Error(), "Error()")
		})
	}
}

func Test_iamServiceError_Error(t *testing.T) {
	type fields struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"invoke", fields{errors.New("test")}, "test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &iamServiceError{
				err: tt.fields.err,
			}
			assert.Equalf(t, tt.want, m.Error(), "Error()")
		})
	}
}

func Test_compareGCPPolicy(t *testing.T) {
	type args struct {
		Roles     *gcpiam.Role
		iacPolicy Sorted
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Pass",
			args{
				Roles:     &gcpiam.Role{IncludedPermissions: []string{"Fred"}},
				iacPolicy: Sorted{GCP: []string{"Fred"}},
			},
			true,
		},
		{
			"False",
			args{
				Roles:     &gcpiam.Role{IncludedPermissions: []string{"Fred"}},
				iacPolicy: Sorted{GCP: []string{"Jane"}},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, compareGCPPolicy(tt.args.Roles, tt.args.iacPolicy), "compareGCPPolicy(%v, %v)", tt.args.Roles, tt.args.iacPolicy)
		})
	}
}

func TestCompareAWSRole_TargetFileAndWriteBaseline(t *testing.T) {
	t.Parallel()

	const directory = "testdata/compare/aws_simple"
	const arn = "arn:aws:iam::123456789012:policy/placeholder"

	t.Run("write then read back round-trips as same, with no live API call", func(t *testing.T) {
		t.Parallel()

		baseline := filepath.Join(t.TempDir(), "baseline.json")

		same, err := compareAWSRole(directory, arn, false, false, baseline, baseline)
		if err != nil {
			t.Fatalf("compareAWSRole() error = %v", err)
		}

		if !same {
			t.Error("compareAWSRole() = false, want true (baseline just written by the same call)")
		}

		if _, err := os.Stat(baseline); err != nil {
			t.Errorf("expected --write-baseline to create %s: %v", baseline, err)
		}
	})

	t.Run("target-file differing from the generated policy reports a difference", func(t *testing.T) {
		t.Parallel()

		stale := filepath.Join(t.TempDir(), "stale.json")
		if err := os.WriteFile(stale, []byte(`{"Version":"2012-10-17","Statement":[{"Sid":"0","Effect":"Allow","Action":["s3:bogus"],"Resource":"*"}]}`), 0o600); err != nil {
			t.Fatal(err)
		}

		same, err := compareAWSRole(directory, arn, false, false, stale, "")
		if err != nil {
			t.Fatalf("compareAWSRole() error = %v", err)
		}

		if same {
			t.Error("compareAWSRole() = true, want false (target-file content deliberately differs)")
		}
	})

	t.Run("missing target-file errors instead of silently falling back to the live API", func(t *testing.T) {
		t.Parallel()

		_, err := compareAWSRole(directory, arn, false, false, filepath.Join(t.TempDir(), "missing.json"), "")
		if err == nil {
			t.Fatal("expected an error for a missing target-file, got nil")
		}

		if _, ok := err.(*readTargetFileError); !ok {
			t.Errorf("expected *readTargetFileError, got %T", err)
		}
	})
}

func TestCompareGCPRole_TargetFileAndWriteBaseline(t *testing.T) {
	t.Parallel()

	const directory = "testdata/compare/gcp_simple"
	const arn = "projects/pike-test/roles/placeholder"

	t.Run("write then read back round-trips as same, with no live API call", func(t *testing.T) {
		t.Parallel()

		baseline := filepath.Join(t.TempDir(), "baseline.json")

		same, err := compareGCPRole(directory, arn, false, false, baseline, baseline)
		if err != nil {
			t.Fatalf("compareGCPRole() error = %v", err)
		}

		if !same {
			t.Error("compareGCPRole() = false, want true (baseline just written by the same call)")
		}
	})

	t.Run("target-file differing from the generated permissions reports a difference", func(t *testing.T) {
		t.Parallel()

		stale := filepath.Join(t.TempDir(), "stale.json")
		if err := os.WriteFile(stale, []byte(`["storage.bogus.permission"]`), 0o600); err != nil {
			t.Fatal(err)
		}

		same, err := compareGCPRole(directory, arn, false, false, stale, "")
		if err != nil {
			t.Fatalf("compareGCPRole() error = %v", err)
		}

		if same {
			t.Error("compareGCPRole() = true, want false (target-file content deliberately differs)")
		}
	})
}

func Test_writeGCPBaseline_readGCPBaseline(t *testing.T) {
	t.Parallel()

	path := filepath.Join(t.TempDir(), "perms.json")
	want := []string{"storage.buckets.create", "storage.buckets.get"}

	if err := writeGCPBaseline(path, []string{"storage.buckets.get", "storage.buckets.create", "storage.buckets.get"}); err != nil {
		t.Fatalf("writeGCPBaseline() error = %v", err)
	}

	got, err := readGCPBaseline(path)
	if err != nil {
		t.Fatalf("readGCPBaseline() error = %v", err)
	}

	assert.Equal(t, want, got)
}

func Test_readGCPBaseline_missingFile(t *testing.T) {
	t.Parallel()

	_, err := readGCPBaseline(filepath.Join(t.TempDir(), "missing.json"))
	if _, ok := err.(*readTargetFileError); !ok {
		t.Errorf("expected *readTargetFileError, got %T (%v)", err, err)
	}
}

func Test_invalidCloudError_Error(t *testing.T) {
	type fields struct {
		arn string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"aws arn", fields{arn: "arn:aws:iam::123456789012:policy/test"}, "Invalid Cloud: arn:aws:iam::123456789012:policy/test"},
		{"invalid format", fields{arn: "invalid-format"}, "Invalid Cloud: invalid-format"},
		{"empty", fields{arn: ""}, "Invalid Cloud: "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &invalidCloudError{
				arn: tt.fields.arn,
			}
			assert.Equalf(t, tt.want, e.Error(), "Error()")
		})
	}
}

func Test_apiNotFoundError_Error(t *testing.T) {
	type fields struct {
		API string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"compute api", fields{API: "compute.googleapis.com"}, "API compute.googleapis.com not found"},
		{"storage api", fields{API: "storage.googleapis.com"}, "API storage.googleapis.com not found"},
		{"empty api", fields{API: ""}, "API  not found"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &apiNotFoundError{
				API: tt.fields.API,
			}
			assert.Equalf(t, tt.want, m.Error(), "Error()")
		})
	}
}

func Test_apiNotEnabledError_Error(t *testing.T) {
	type fields struct {
		API string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"compute api", fields{API: "compute.googleapis.com"}, "API compute.googleapis.com not enabled"},
		{"storage api", fields{API: "storage.googleapis.com"}, "API storage.googleapis.com not enabled"},
		{"empty api", fields{API: ""}, "API  not enabled"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &apiNotEnabledError{
				API: tt.fields.API,
			}
			assert.Equalf(t, tt.want, m.Error(), "Error()")
		})
	}
}

func Test_formatterError_Error(t *testing.T) {
	type fields struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"format error", fields{err: errors.New("failed to format")}, "formatter failed: failed to format"},
		{"nil error", fields{err: nil}, "formatter failed: <nil>"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &formatterError{
				err: tt.fields.err,
			}
			assert.Equalf(t, tt.want, m.Error(), "Error()")
		})
	}
}

func Test_arnNotVerified_Error(t *testing.T) {
	type fields struct {
		arn string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"aws arn", fields{arn: "arn:aws:iam::123456789012:policy/test"}, "ARN arn:aws:iam::123456789012:policy/test not verified: expected format arn:aws:iam::123456789012:role/role-name"},
		{"invalid arn", fields{arn: "invalid"}, "ARN invalid not verified: expected format arn:aws:iam::123456789012:role/role-name"},
		{"empty arn", fields{arn: ""}, "ARN  not verified: expected format arn:aws:iam::123456789012:role/role-name"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &arnNotVerified{
				arn: tt.fields.arn,
			}
			assert.Equalf(t, tt.want, e.Error(), "Error()")
		})
	}
}
