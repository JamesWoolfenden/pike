//go:build auth
// +build auth

package pike_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/google/go-github/v47/github"
	pike "github.com/jameswoolfenden/pike/src"
	"golang.org/x/oauth2"
)

func TestSetRepoSecret(t *testing.T) {
	t.Parallel()

	type args struct {
		repository string
		keyText    string
		keyName    string
	}

	tests := []struct {
		name    string
		args    args
		want    *github.Response
		wantErr bool
	}{
		{
			name: "invalid repository format",
			args: args{
				repository: "invalid-repo-format",
				keyText:    "secret-value",
				keyName:    "TEST_SECRET",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "valid repository format",
			args: args{
				repository: "owner/repo",
				keyText:    "test-secret-value",
				keyName:    "TEST_SECRET",
			},
			want:    &github.Response{},
			wantErr: false,
		},
		{
			name: "empty secret value",
			args: args{
				repository: "owner/repo",
				keyText:    "",
				keyName:    "TEST_SECRET",
			},
			want:    &github.Response{},
			wantErr: false,
		},
		{
			name: "empty secret name",
			args: args{
				repository: "owner/repo",
				keyText:    "test-secret-value",
				keyName:    "",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "special characters in repository",
			args: args{
				repository: "test-owner/test-repo-123",
				keyText:    "test-secret-value",
				keyName:    "TEST_SECRET_123",
			},
			want:    &github.Response{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := pike.SetRepoSecret(tt.args.repository, tt.args.keyText, tt.args.keyName)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetRepoSecret() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetRepoSecret() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getGithubClient(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		envVars map[string]string
		wantNil bool
	}{
		{
			name: "with GITHUB_TOKEN",
			envVars: map[string]string{
				"GITHUB_TOKEN": "test-token",
			},
			wantNil: false,
		},
		{
			name: "with GITHUB_API",
			envVars: map[string]string{
				"GITHUB_API": "test-api-token",
			},
			wantNil: false,
		},
		{
			name:    "no token set",
			envVars: map[string]string{},
			wantNil: false,
		},
		{
			name: "empty GITHUB_TOKEN",
			envVars: map[string]string{
				"GITHUB_TOKEN": "",
			},
			wantNil: false,
		},
		{
			name: "both tokens set",
			envVars: map[string]string{
				"GITHUB_TOKEN": "test-token",
				"GITHUB_API":   "test-api-token",
			},
			wantNil: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Clear environment variables
			os.Clearenv()

			// Set test environment variables
			for k, v := range tt.envVars {
				err := os.Setenv(k, v)
				if err != nil {
					return
				}
			}

			ctx, client := pike.GetGithubClient()

			if ctx == nil {
				t.Error("GetGithubClient() returned nil context")
			}

			if (client == nil) != tt.wantNil {
				t.Errorf("GetGithubClient() returned nil client = %v, want nil = %v", client == nil, tt.wantNil)
			}

			// Verify client has transport configured when token is provided
			if len(tt.envVars) > 0 && client != nil {
				transport, ok := client.Client().Transport.(*oauth2.Transport)
				if !ok {
					t.Error("GetGithubClient() client does not have oauth2 transport")
				}
				if transport == nil {
					t.Error("GetGithubClient() client has nil transport")
				}
			}
		})
	}
}

func TestRemote_New(t *testing.T) {
	t.Parallel()

	type args struct {
		target     string
		repository string
		region     string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid inputs",
			args: args{
				target:     "arn:aws:iam::123456789012:role/test-role",
				repository: "owner/repo",
				region:     "us-west-2",
			},
			wantErr: false,
		},
		{
			name: "invalid target",
			args: args{
				target:     "invalid-arn",
				repository: "owner/repo",
				region:     "us-west-2",
			},
			wantErr: true,
		},
		{
			name: "empty repository",
			args: args{
				target:     "arn:aws:iam::123456789012:role/test-role",
				repository: "",
				region:     "us-west-2",
			},
			wantErr: true,
		},
		{
			name: "invalid region",
			args: args{
				target:     "arn:aws:iam::123456789012:role/test-role",
				repository: "owner/repo",
				region:     "invalid-region",
			},
			wantErr: true,
		},
		{
			name: "empty region",
			args: args{
				target:     "arn:aws:iam::123456789012:role/test-role",
				repository: "owner/repo",
				region:     "",
			},
			wantErr: true,
		},
		{
			name: "malformed repository",
			args: args{
				target:     "arn:aws:iam::123456789012:role/test-role",
				repository: "invalid-format",
				region:     "us-west-2",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := pike.Remote(tt.args.target, tt.args.repository, tt.args.region); (err != nil) != tt.wantErr {
				t.Errorf("Remote() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetPublicKeyDetails_Integration(t *testing.T) {
	t.Parallel()

	type args struct {
		owner      string
		repository string
	}

	tests := []struct {
		name        string
		args        args
		wantKeyID   string
		wantPkValue string
		wantErr     bool
	}{
		{
			name: "valid repository",
			args: args{
				owner:      "jameswoolfenden",
				repository: "pike",
			},
			wantKeyID:   "test-key-id",
			wantPkValue: "test-public-key",
			wantErr:     false,
		},
		{
			name: "invalid repository",
			args: args{
				owner:      "nonexistent",
				repository: "nonexistent",
			},
			wantKeyID:   "",
			wantPkValue: "",
			wantErr:     true,
		},
		{
			name: "empty owner",
			args: args{
				owner:      "",
				repository: "pike",
			},
			wantKeyID:   "",
			wantPkValue: "",
			wantErr:     true,
		},
		{
			name: "empty repository",
			args: args{
				owner:      "jameswoolfenden",
				repository: "",
			},
			wantKeyID:   "",
			wantPkValue: "",
			wantErr:     true,
		},
		{
			name: "special characters",
			args: args{
				owner:      "test@#$%",
				repository: "repo@#$%",
			},
			wantKeyID:   "",
			wantPkValue: "",
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotKeyID, gotPkValue, err := pike.GetPublicKeyDetails(tt.args.owner, tt.args.repository)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPublicKeyDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && gotKeyID == "" {
				t.Error("GetPublicKeyDetails() expected non-empty keyID for successful case")
			}
			if !tt.wantErr && gotPkValue == "" {
				t.Error("GetPublicKeyDetails() expected non-empty pkValue for successful case")
			}
		})
	}
}
