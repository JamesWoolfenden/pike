//go:build auth

package pike

import (
	"os"
	"path/filepath"
	"testing"

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

func TestCompare(t *testing.T) {
	t.Parallel()

	type args struct {
		directory string
		arn       string
		init      bool
	}

	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"fail arn is empty", args{"./testdata/init/nicconf", "", false}, false, true},
		{"fail arn is not policy", args{"./testdata/init/nicconf", "arn:aws:iam::680235478471:user/readonly", false}, false, true},
		{"pass", args{"./testdata/init/nicconf", "arn:aws:iam::680235478471:policy/testdata", false}, true, false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := Compare(tt.args.directory, tt.args.arn, tt.args.init)
			if (err != nil) != tt.wantErr {
				t.Errorf("Compare() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Compare() got = %v, want %v", got, tt.want)
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
			wantErr:   &invalidARNError{},
		},
		{
			name:      "valid inputs",
			directory: tmpDir,
			arn:       "arn:aws:iam::123456789012:policy/test",
			wantBool:  false,
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
