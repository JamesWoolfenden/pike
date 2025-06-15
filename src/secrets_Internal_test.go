package pike

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

func Test_splitHub(t *testing.T) {
	t.Parallel()

	type args struct {
		repository string
	}

	tests := []struct {
		name    string
		args    args
		want    string
		want1   string
		wantErr bool
	}{
		{
			name: "valid short format",
			args: args{
				repository: "jameswoolfenden/pike",
			},
			want:    "jameswoolfenden",
			want1:   "pike",
			wantErr: false,
		},
		{
			name: "valid long format",
			args: args{
				repository: "https://github.com/jameswoolfenden/pike",
			},
			want:    "jameswoolfenden",
			want1:   "pike",
			wantErr: false,
		},
		{
			name: "invalid format",
			args: args{
				repository: "jameswoolfenden/pike/extra",
			},
			want:    "",
			want1:   "",
			wantErr: true,
		},
		{
			name: "empty string",
			args: args{
				repository: "",
			},
			want:    "",
			want1:   "",
			wantErr: true,
		},
		{
			name: "single segment",
			args: args{
				repository: "onlyone",
			},
			want:    "",
			want1:   "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, got1, err := SplitHub(tt.args.repository)
			if (err != nil) != tt.wantErr {
				t.Errorf("SplitHub() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SplitHub() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SplitHub() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestEncryptPlaintext_Extended(t *testing.T) {
	t.Parallel()

	type args struct {
		plaintext    string
		publicKeyB64 string
	}

	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "valid encryption",
			args: args{
				plaintext:    "test secret",
				publicKeyB64: "VGhpcyBpcyBhIHZhbGlkIGJhc2U2NCBlbmNvZGVkIHB1YmxpYyBrZXk=",
			},
			wantErr: false,
		},
		{
			name: "empty plaintext",
			args: args{
				plaintext:    "",
				publicKeyB64: "VGhpcyBpcyBhIHZhbGlkIGJhc2U2NCBlbmNvZGVkIHB1YmxpYyBrZXk=",
			},
			wantErr: false,
		},
		{
			name: "invalid base64 public key",
			args: args{
				plaintext:    "test secret",
				publicKeyB64: "invalid-base64!@#$",
			},
			wantErr: true,
		},
		{
			name: "empty public key",
			args: args{
				plaintext:    "test secret",
				publicKeyB64: "",
			},
			wantErr: true,
		},
		//{
		//	name: "public key too short",
		//	args: args{
		//		plaintext:    "test secret",
		//		publicKeyB64: "aGVsbG8=",
		//	},
		//	wantErr: true,
		//},
		{
			name: "very long plaintext",
			args: args{
				plaintext:    string(make([]byte, 1024*1024)), // 1MB of data
				publicKeyB64: "VGhpcyBpcyBhIHZhbGlkIGJhc2U2NCBlbmNvZGVkIHB1YmxpYyBrZXk=",
			},
			wantErr: false,
		},
		{
			name: "special characters in plaintext",
			args: args{
				plaintext:    "!@#$%^&*()_+{}|:<>?~`-=[]\\;',./",
				publicKeyB64: "VGhpcyBpcyBhIHZhbGlkIGJhc2U2NCBlbmNvZGVkIHB1YmxpYyBrZXk=",
			},
			wantErr: false,
		},
		{
			name: "unicode characters in plaintext",
			args: args{
				plaintext:    "Hello 世界 🌍",
				publicKeyB64: "VGhpcyBpcyBhIHZhbGlkIGJhc2U2NCBlbmNvZGVkIHB1YmxpYyBrZXk=",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := EncryptPlaintext(tt.args.plaintext, tt.args.publicKeyB64)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncryptPlaintext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if got == nil {
					t.Error("EncryptPlaintext() returned nil for successful encryption")
				}
				if len(got) == 0 {
					t.Error("EncryptPlaintext() returned empty bytes for successful encryption")
				}
			}
		})
	}
}

func TestAwsCredentialsError_Error(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected string
	}{
		{
			name:     "simple error message",
			err:      errors.New("connection timeout"),
			expected: "failed to get AWS credentials: connection timeout",
		},
		{
			name:     "empty error message",
			err:      errors.New(""),
			expected: "failed to get AWS credentials: ",
		},
		{
			name:     "formatted error message",
			err:      fmt.Errorf("invalid region: %s", "us-invalid-1"),
			expected: "failed to get AWS credentials: invalid region: us-invalid-1",
		},
		{
			name:     "wrapped error",
			err:      fmt.Errorf("wrapped: %w", errors.New("original error")),
			expected: "failed to get AWS credentials: wrapped: original error",
		},
		{
			name:     "nil error",
			err:      nil,
			expected: "failed to get AWS credentials: <nil>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &awsCredentialsError{err: tt.err}
			result := e.Error()

			if result != tt.expected {
				t.Errorf("awsCredentialsError.Error() = %q, expected %q", result, tt.expected)
			}
		})
	}
}

func TestAwsCredentialsError_ErrorInterface(t *testing.T) {
	err := &awsCredentialsError{err: errors.New("test error")}

	// Verify it implements the error interface
	var _ error = err

	// Verify it can be used as an error
	if err.Error() == "" {
		t.Error("awsCredentialsError should return non-empty error message")
	}
}

func TestAwsCredentialsError_ErrorFormatting(t *testing.T) {
	testErr := errors.New("access denied")
	awsErr := &awsCredentialsError{err: testErr}

	errorMsg := awsErr.Error()

	// Verify the error message contains the expected prefix
	expectedPrefix := "failed to get AWS credentials:"
	if !strings.HasPrefix(errorMsg, expectedPrefix) {
		t.Errorf("Error message should start with %q, got %q", expectedPrefix, errorMsg)
	}

	// Verify the original error message is included
	if !strings.Contains(errorMsg, testErr.Error()) {
		t.Errorf("Error message should contain original error %q, got %q", testErr.Error(), errorMsg)
	}
}

func TestAwsCredentialsError_MultipleInstances(t *testing.T) {
	err1 := &awsCredentialsError{err: errors.New("error 1")}
	err2 := &awsCredentialsError{err: errors.New("error 2")}

	msg1 := err1.Error()
	msg2 := err2.Error()

	if msg1 == msg2 {
		t.Error("Different awsCredentialsError instances should produce different error messages")
	}

	if !strings.Contains(msg1, "error 1") {
		t.Errorf("First error should contain 'error 1', got %q", msg1)
	}

	if !strings.Contains(msg2, "error 2") {
		t.Errorf("Second error should contain 'error 2', got %q", msg2)
	}
}
