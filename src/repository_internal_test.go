package pike

import (
	"errors"
	"testing"
)

func TestGitCloneError_Error(t *testing.T) {
	tests := []struct {
		name        string
		repository  string
		destination string
		err         error
		expected    string
	}{
		{
			name:        "basic error message",
			repository:  "https://github.com/user/repo.git",
			destination: "/tmp/repo",
			err:         errors.New("connection failed"),
			expected:    "failed to clone repository https://github.com/user/repo.git /tmp/repo connection failed",
		},
		{
			name:        "empty repository",
			repository:  "",
			destination: "/tmp/repo",
			err:         errors.New("invalid URL"),
			expected:    "failed to clone repository  /tmp/repo invalid URL",
		},
		{
			name:        "empty destination",
			repository:  "https://github.com/user/repo.git",
			destination: "",
			err:         errors.New("no destination"),
			expected:    "failed to clone repository https://github.com/user/repo.git  no destination",
		},
		{
			name:        "nil error",
			repository:  "https://github.com/user/repo.git",
			destination: "/tmp/repo",
			err:         nil,
			expected:    "failed to clone repository https://github.com/user/repo.git /tmp/repo <nil>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &gitCloneError{
				repository:  tt.repository,
				destination: tt.destination,
				err:         tt.err,
			}
			if got := e.Error(); got != tt.expected {
				t.Errorf("gitCloneError.Error() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestGitHeadError_Error(t *testing.T) {
	tests := []struct {
		name        string
		repository  string
		destination string
		err         error
		expected    string
	}{
		{
			name:        "basic error message",
			repository:  "https://github.com/user/repo.git",
			destination: "/tmp/repo",
			err:         errors.New("head not found"),
			expected:    "failed to get head https://github.com/user/repo.git /tmp/repo head not found",
		},
		{
			name:        "empty strings",
			repository:  "",
			destination: "",
			err:         errors.New("reference error"),
			expected:    "failed to get head   reference error",
		},
		{
			name:        "nil error",
			repository:  "repo",
			destination: "dest",
			err:         nil,
			expected:    "failed to get head repo dest <nil>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &gitHeadError{
				repository:  tt.repository,
				destination: tt.destination,
				err:         tt.err,
			}
			if got := e.Error(); got != tt.expected {
				t.Errorf("gitHeadError.Error() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestGitCommitObjectError_Error(t *testing.T) {
	tests := []struct {
		name        string
		repository  string
		destination string
		err         error
		expected    string
	}{
		{
			name:        "basic error message",
			repository:  "https://github.com/user/repo.git",
			destination: "/tmp/repo",
			err:         errors.New("commit not found"),
			expected:    "failed to get commit object https://github.com/user/repo.git /tmp/repo commit not found",
		},
		{
			name:        "long repository URL",
			repository:  "https://very-long-domain-name.example.com/organization/very-long-repository-name.git",
			destination: "/very/long/path/to/destination/directory",
			err:         errors.New("object does not exist"),
			expected:    "failed to get commit object https://very-long-domain-name.example.com/organization/very-long-repository-name.git /very/long/path/to/destination/directory object does not exist",
		},
		{
			name:        "all empty except error",
			repository:  "",
			destination: "",
			err:         errors.New("some error"),
			expected:    "failed to get commit object   some error",
		},
		{
			name:        "nil error",
			repository:  "repo",
			destination: "dest",
			err:         nil,
			expected:    "failed to get commit object repo dest <nil>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &gitCommitObjectError{
				repository:  tt.repository,
				destination: tt.destination,
				err:         tt.err,
			}
			if got := e.Error(); got != tt.expected {
				t.Errorf("gitCommitObjectError.Error() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestErrorTypes_ImplementErrorInterface(t *testing.T) {
	var err error

	// Test that all error types implement the error interface
	err = &gitCloneError{repository: "repo", destination: "dest", err: errors.New("test")}
	if err.Error() == "" {
		t.Error("gitCloneError should implement error interface")
	}

	err = &gitHeadError{repository: "repo", destination: "dest", err: errors.New("test")}
	if err.Error() == "" {
		t.Error("gitHeadError should implement error interface")
	}

	err = &gitCommitObjectError{repository: "repo", destination: "dest", err: errors.New("test")}
	if err.Error() == "" {
		t.Error("gitCommitObjectError should implement error interface")
	}
}
