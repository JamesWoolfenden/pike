package pike

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	diff "github.com/yudai/gojsondiff"
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
		//{
		//	name:        "Valid policy and diff",
		//	policy:      `{"Version": "2012-10-17", "Statement": [{"Effect": "Allow"}]}`,
		//	diff:        &mockDiff{},
		//	wantBool:    false,
		//	wantErr:     false,
		//	description: "Should successfully format and display differences",
		//},
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
			assert.Equal(t, tt.expected, *result)
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
