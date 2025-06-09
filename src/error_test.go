package pike

import (
	"errors"
	"fmt"
	"testing"
)

func Test_notImplementedResourceError_Error(t *testing.T) {
	t.Parallel()

	type fields struct {
		Name string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Pass", fields{Name: "test"}, "not implemented test"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			m := &notImplementedResourceError{
				Name: tt.fields.Name,
			}

			if got := m.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_notImplementedDatasourceError_Error(t *testing.T) {
	t.Parallel()

	type fields struct {
		Name string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"pass", fields{Name: "test"}, "data.test not implemented"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			m := &notImplementedDatasourceError{
				Name: tt.fields.Name,
			}

			if got := m.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unknownPermissionError_Error(t *testing.T) {
	t.Parallel()

	type fields struct {
		Name string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"pass", fields{Name: "test"}, "unknown permission resource type test"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			m := &unknownPermissionError{
				Name: tt.fields.Name,
			}

			if got := m.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGitReferenceError_Error(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		m    *gitReferenceError
		want string
	}{
		{
			name: "basic git reference",
			m: &gitReferenceError{
				name: "refs/heads/main",
			},
			want: "git reference in module source path unsupported refs/heads/main",
		},
		{
			name: "empty reference",
			m: &gitReferenceError{
				name: "",
			},
			want: "git reference in module source path unsupported ",
		},
		{
			name: "commit hash reference",
			m: &gitReferenceError{
				name: "a1b2c3d4e5f6",
			},
			want: "git reference in module source path unsupported a1b2c3d4e5f6",
		},
		{
			name: "tag reference",
			m: &gitReferenceError{
				name: "refs/tags/v1.0.0",
			},
			want: "git reference in module source path unsupported refs/tags/v1.0.0",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.m.Error(); got != tt.want {
				t.Errorf("gitReferenceError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBackendExistsError_Error(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		m    *backendExistsError
		want string
	}{
		{
			name: "basic backend error",
			m:    &backendExistsError{},
			want: "backend already exists",
		},
		{
			name: "new instance",
			m:    new(backendExistsError),
			want: "backend already exists",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.m.Error(); got != tt.want {
				t.Errorf("backendExistsError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepositoryFormatError_Error(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		m    *repositoryFormatError
		want string
	}{
		{
			name: "empty repository",
			m: &repositoryFormatError{
				name: "",
			},
			want: "repository not formatted correctly ",
		},
		{
			name: "invalid path format",
			m: &repositoryFormatError{
				name: "invalid/repo/path///",
			},
			want: "repository not formatted correctly invalid/repo/path///",
		},
		{
			name: "special characters",
			m: &repositoryFormatError{
				name: "@#$%^&*",
			},
			want: "repository not formatted correctly @#$%^&*",
		},
		{
			name: "url format",
			m: &repositoryFormatError{
				name: "https://github.com/user/repo.git",
			},
			want: "repository not formatted correctly https://github.com/user/repo.git",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.m.Error(); got != tt.want {
				t.Errorf("repositoryFormatError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepositoryFormatError_Error_Additional(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		m    *repositoryFormatError
		want string
	}{
		{
			name: "whitespace only",
			m: &repositoryFormatError{
				name: "   ",
			},
			want: "repository not formatted correctly    ",
		},
		{
			name: "very long path",
			m: &repositoryFormatError{
				name: "org/extremely/long/repository/path/that/exceeds/normal/length/limits/and/tests/boundary/conditions",
			},
			want: "repository not formatted correctly org/extremely/long/repository/path/that/exceeds/normal/length/limits/and/tests/boundary/conditions",
		},
		{
			name: "unicode characters",
			m: &repositoryFormatError{
				name: "организация/репозиторий",
			},
			want: "repository not formatted correctly организация/репозиторий",
		},
		{
			name: "numeric repository",
			m: &repositoryFormatError{
				name: "12345/67890",
			},
			want: "repository not formatted correctly 12345/67890",
		},
		{
			name: "mixed case with symbols",
			m: &repositoryFormatError{
				name: "My-Org/Repo_Name-123",
			},
			want: "repository not formatted correctly My-Org/Repo_Name-123",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.m.Error(); got != tt.want {
				t.Errorf("repositoryFormatError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMappingsEmpty_Error(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		m    *mappingsEmptyError
		want string
	}{
		{
			name: "basic empty mappings",
			m:    &mappingsEmptyError{},
			want: "mappings are empty",
		},
		{
			name: "new instance",
			m:    new(mappingsEmptyError),
			want: "mappings are empty",
		},
		{
			name: "nil pointer",
			m:    (*mappingsEmptyError)(nil),
			want: "mappings are empty",
		},
		{
			name: "multiple calls",
			m:    &mappingsEmptyError{},
			want: "mappings are empty",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.m.Error(); got != tt.want {
				t.Errorf("mappingsEmptyError.Error() = %v, want %v", got, tt.want)
			}

			// Test multiple calls return the same result
			if tt.name == "multiple calls" {
				for i := 0; i < 3; i++ {
					if got := tt.m.Error(); got != tt.want {
						t.Errorf("mappingsEmptyError.Error() iteration %d = %v, want %v", i, got, tt.want)
					}
				}
			}
		})
	}
}

func TestEmptyPermissionsError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		want string
	}{
		{
			name: "returns correct error message",
			want: "permissions list cannot be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := &emptyPermissionsError{}

			if got := err.Error(); got != tt.want {
				t.Errorf("emptyPermissionsError.Error() = %v, want %v", got, tt.want)
			}
		})
	}

	// Verify it implements error interface
	var _ error = &emptyPermissionsError{}
}

func TestEmptyTypeNameError(t *testing.T) {
	t.Parallel()
	t.Run("returns correct error message", func(t *testing.T) {
		t.Parallel()

		err := &emptyTypeNameError{}

		expected := "TypeName cannot be empty"

		if got := err.Error(); got != expected {
			t.Errorf("emptyTypeNameError.Error() = %v, want %v", got, expected)
		}
	})

	//goland:noinspection GoLinter
	t.Run("implements error interface", func(t *testing.T) {
		t.Parallel()
		var err error = &emptyTypeNameError{} // Verify it satisfies error interface

		//goland:noinspection GoLinter
		if err == nil {
			t.Error("emptyTypeNameError should implement error interface")
		}
	})
}

func TestEmptyNameError(t *testing.T) {
	t.Parallel()

	err := &emptyNameError{}

	expected := "Name cannot be empty"
	if got := err.Error(); got != expected {
		t.Errorf("emptyNameError.Error() = %v, want %v", got, expected)
	}
}

func TestEmptyNameError_ImplementsError(t *testing.T) {
	var _ error = (*emptyNameError)(nil)
}

func TestAssertionFailedError(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		message  string
		err      error
		expected string
	}{
		{
			name:     "basic error message",
			message:  "test failed",
			err:      errors.New("invalid input"),
			expected: "assertion failed: test failed invalid input",
		},
		{
			name:     "empty message",
			message:  "",
			err:      errors.New("error only"),
			expected: "assertion failed:  error only",
		},
		{
			name:     "nil error",
			message:  "test message",
			err:      nil,
			expected: "assertion failed: test message <nil>",
		},
		{
			name:     "empty message and nil error",
			message:  "",
			err:      nil,
			expected: "assertion failed:  <nil>",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			err := &assertionFailedError{
				message: tc.message,
				err:     tc.err,
			}

			if got := err.Error(); got != tc.expected {
				t.Errorf("assertionFailedError.Error() = %v, want %v", got, tc.expected)
			}
		})
	}
}

func TestInvalidJSONError(t *testing.T) {
	t.Parallel()
	t.Run("implements error interface", func(t *testing.T) {
		t.Parallel()

		var _ error = &invalidJSONError{}
	})

	t.Run("returns correct error message", func(t *testing.T) {
		t.Parallel()

		err := &invalidJSONError{}
		expected := "invalid json, was empty or corrupt"

		if got := err.Error(); got != expected {
			t.Errorf("invalidJSONError.Error() = %v, want %v", got, expected)
		}
	})
}

func TestUnmarshallJSONError(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		err      error
		resource string
		want     string
	}{
		{
			name:     "basic error",
			err:      errors.New("parse error"),
			resource: "test-resource",
			want:     "failed to unmarshal json parse error for test-resource",
		},
		{
			name:     "nil error",
			err:      nil,
			resource: "test-resource",
			want:     "failed to unmarshal json <nil> for test-resource",
		},
		{
			name:     "empty resource",
			err:      errors.New("parse error"),
			resource: "",
			want:     "failed to unmarshal json parse error for ",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			err := &unmarshallJSONError{
				err:      tc.err,
				resource: tc.resource,
			}

			if got := err.Error(); got != tc.want {
				t.Errorf("unmarshallJSONError.Error() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestAttributesFieldMissingError(t *testing.T) {
	err := &attributesFieldMissingError{}
	want := "attributes field missing"

	if got := err.Error(); got != want {
		t.Errorf("attributesFieldMissingError.Error() = %v, want %v", got, want)
	}
}

func TestAssertionError(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		message string
		want    string
	}{
		{
			name:    "basic message",
			message: "test failed",
			want:    "assertion failed for: test failed",
		},
		{
			name:    "empty message",
			message: "",
			want:    "assertion failed for: ",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			err := &assertionError{message: tc.message}

			if got := err.Error(); got != tc.want {
				t.Errorf("assertionError.Error() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestTemplateParseError(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		err  error
		want string
	}{
		{
			name: "with error",
			err:  errors.New("invalid syntax"),
			want: "failed to parse template invalid syntax",
		},
		{
			name: "nil error",
			err:  nil,
			want: "failed to parse template <nil>",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			err := &templateParseError{err: tc.err}

			if got := err.Error(); got != tc.want {
				t.Errorf("templateParseError.Error() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestTemplateExecuteError(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		err  error
		want string
	}{
		{
			name: "with error",
			err:  errors.New("execution failed"),
			want: "failed to execute template execution failed",
		},
		{
			name: "nil error",
			err:  nil,
			want: "failed to execute template <nil>",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			err := &templateExecuteError{err: tc.err}

			if got := err.Error(); got != tc.want {
				t.Errorf("templateExecuteError.Error() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestMarshallAWSPolicyError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		err      error
		expected string
	}{
		{
			name:     "with nil error",
			err:      nil,
			expected: "failed to marshal policy: <nil>",
		},
		{
			name:     "with simple error",
			err:      errors.New("invalid format"),
			expected: "failed to marshal policy: invalid format",
		},
		{
			name:     "with wrapped error",
			err:      fmt.Errorf("wrapped: %w", errors.New("base error")),
			expected: "failed to marshal policy: wrapped: base error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := &marshallAWSPolicyError{err: tt.err}

			if got := err.Error(); got != tt.expected {
				t.Errorf("marshallAWSPolicyError.Error() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestDirectoryErrors(t *testing.T) {
	t.Parallel()
	t.Run("empty directory error", func(t *testing.T) {
		t.Parallel()

		err := &emptyDirectoryError{}
		expected := "directory value cannot be an empty string"

		if err.Error() != expected {
			t.Errorf("expected %q, got %q", expected, err.Error())
		}
	})

	t.Run("directory not found error", func(t *testing.T) {
		t.Parallel()
		err := &directoryNotFoundError{directory: "/tmp/nonexistent"}
		expected := "directory does not exist: /tmp/nonexistent"

		if err.Error() != expected {
			t.Errorf("expected %q, got %q", expected, err.Error())
		}
	})
}

func TestARNErrors(t *testing.T) {
	t.Run("empty ARN error", func(t *testing.T) {
		t.Parallel()
		err := &arnEmptyError{}
		expected := "ARN cannot be empty"

		if err.Error() != expected {
			t.Errorf("expected %q, got %q", expected, err.Error())
		}
	})

	t.Run("invalid ARN error", func(t *testing.T) {
		t.Parallel()

		err := &invalidARNError{arn: "invalid:arn"}

		expected := "invalid role or ARN: invalid:arn"

		if err.Error() != expected {
			t.Errorf("expected %q, got %q", expected, err.Error())
		}
	})
}

func TestAWSErrors(t *testing.T) {
	t.Parallel()
	t.Run("AWS config error", func(t *testing.T) {
		t.Parallel()

		err := &awsConfigError{err: &emptyNameError{}}

		expected := "failed to load AWS config: Name cannot be empty"

		if err.Error() != expected {
			t.Errorf("expected %q, got %q", expected, err.Error())
		}
	})

	t.Run("get IAM version error", func(t *testing.T) {
		t.Parallel()

		err := &getIAMVersionError{err: &emptyNameError{}}

		expected := "failed to get IAM version: Name cannot be empty"

		if err.Error() != expected {
			t.Errorf("expected %q, got %q", expected, err.Error())
		}
	})
}

func TestPolicyErrors(t *testing.T) {
	t.Parallel()
	t.Run("sort actions error", func(t *testing.T) {
		t.Parallel()

		err := &sortActionsError{json: "invalid-json"}
		expected := "failed to sort actions: invalid-json"

		if err.Error() != expected {
			t.Errorf("expected %q, got %q", expected, err.Error())
		}
	})

	t.Run("get policy version error", func(t *testing.T) {
		t.Parallel()

		err := &getPolicyVersionError{err: &emptyNameError{}}
		expected := "failed to get policy version: Name cannot be empty"

		if err.Error() != expected {
			t.Errorf("expected %q, got %q", expected, err.Error())
		}
	})

	t.Run("input validation error", func(t *testing.T) {
		t.Parallel()

		err := &inputValidationError{err: &emptyNameError{}}
		expected := "input validation failed: Name cannot be empty"

		if err.Error() != expected {
			t.Errorf("expected %q, got %q", expected, err.Error())
		}
	})

	t.Run("marshall policy error", func(t *testing.T) {
		t.Parallel()

		err := &marshallPolicyError{err: &emptyNameError{}}

		expected := "failed to marshal policy: Name cannot be empty"

		if err.Error() != expected {
			t.Errorf("expected %q, got %q", expected, err.Error())
		}
	})
}

func TestTerraformErrors(t *testing.T) {
	testErr := errors.New("test error")

	t.Run("terraformPlanError", func(t *testing.T) {
		t.Parallel()

		err := &terraformPlanError{err: testErr}
		expected := "failed to plan terraform test error"

		if err.Error() != expected {
			t.Errorf("got %q, want %q", err.Error(), expected)
		}
	})

	t.Run("terraformNewError", func(t *testing.T) {
		t.Parallel()

		err := &terraformNewError{err: testErr}
		expected := "failed to create terraform test error"

		if err.Error() != expected {
			t.Errorf("got %q, want %q", err.Error(), expected)
		}
	})

	t.Run("terraformOutputError", func(t *testing.T) {
		t.Parallel()

		err := &terraformOutputError{}
		expected := "terraform output is empty"

		if err.Error() != expected {
			t.Errorf("got %q, want %q", err.Error(), expected)
		}
	})

	t.Run("terraformApplyError with target", func(t *testing.T) {
		t.Parallel()

		err := &terraformApplyError{target: "module.test", err: testErr}
		expected := "failed to apply terraform module.test test error"

		if err.Error() != expected {
			t.Errorf("got %q, want %q", err.Error(), expected)
		}
	})

	t.Run("terraformApplyError without target", func(t *testing.T) {
		err := &terraformApplyError{target: "", err: testErr}
		expected := "failed to apply terraform test error"
		if err.Error() != expected {
			t.Errorf("got %q, want %q", err.Error(), expected)
		}
	})
}

func TestSecretAndEncryptionErrors(t *testing.T) {
	testErr := errors.New("test error")

	t.Run("getPublicKeyDetailsError", func(t *testing.T) {
		err := &getPublicKeyDetailsError{err: testErr}
		expected := "failed to get public key details: test error"

		if err.Error() != expected {
			t.Errorf("got %q, want %q", err.Error(), expected)
		}
	})

	t.Run("updateSecretError", func(t *testing.T) {
		err := &updateSecretError{err: testErr}
		expected := "failed to update secret: test error"

		if err.Error() != expected {
			t.Errorf("got %q, want %q", err.Error(), expected)
		}
	})

	t.Run("decodeStringError", func(t *testing.T) {
		t.Parallel()

		err := &decodeStringError{err: testErr}
		expected := "failed to decode string: test error"

		if err.Error() != expected {
			t.Errorf("got %q, want %q", err.Error(), expected)
		}
	})

	t.Run("encryptPlaintextError", func(t *testing.T) {
		t.Parallel()

		err := &encryptPlaintextError{err: testErr}
		expected := "failed to encrypt plaintext: test error"

		if err.Error() != expected {
			t.Errorf("got %q, want %q", err.Error(), expected)
		}
	})

	t.Run("emptyKeyError", func(t *testing.T) {
		err := &emptyKeyError{}
		expected := "empty key"

		if err.Error() != expected {
			t.Errorf("got %q, want %q", err.Error(), expected)
		}
	})

	t.Run("encryptError", func(t *testing.T) {
		err := &encryptError{err: testErr}
		expected := "failed to encrypt: test error"

		if err.Error() != expected {
			t.Errorf("got %q, want %q", err.Error(), expected)
		}
	})
}

func TestGetAWSDataPermissionsError(t *testing.T) {
	originalErr := errors.New("original error")
	customErr := &getAWSDataPermissionsError{err: originalErr}
	expected := "failed to get AWS data permissions original error"

	if customErr.Error() != expected {
		t.Errorf("Expected %s, but got %s", expected, customErr.Error())
	}
}

func TestSplitHubError(t *testing.T) {
	t.Parallel()

	originalErr := errors.New("another error")
	customErr := &splitHubError{err: originalErr}
	expected := "failed to split hub: another error"

	if customErr.Error() != expected {
		t.Errorf("Expected %s, but got %s", expected, customErr.Error())
	}
}

func TestSetRepoSecretError(t *testing.T) {
	originalErr := errors.New("some error")
	customErr := &setRepoSecretError{repository: "my-repo", err: originalErr}
	expected := "failed to set repo secret:my-repo some error"

	if customErr.Error() != expected {
		t.Errorf("Expected %s, but got %s", expected, customErr.Error())
	}
}

func TestSetAWSAuthError(t *testing.T) {
	originalErr := errors.New("auth error")
	customErr := &setAWSAuthError{err: originalErr}
	expected := "failed to set AWS auth error"

	if customErr.Error() != expected {
		t.Errorf("Expected %s, but got %s", expected, customErr.Error())
	}
}

func TestMakeRoleError(t *testing.T) {
	originalErr := errors.New("role error")
	customErr := &makeRoleError{err: originalErr}
	expected := "failed to make role: role error"

	if customErr.Error() != expected {
		t.Errorf("Expected %s, but got %s", expected, customErr.Error())
	}
}
