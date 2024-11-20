package pike

import "testing"

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
				Name: "refs/heads/main",
			},
			want: "git reference in module source path unsupported refs/heads/main",
		},
		{
			name: "empty reference",
			m: &gitReferenceError{
				Name: "",
			},
			want: "git reference in module source path unsupported ",
		},
		{
			name: "commit hash reference",
			m: &gitReferenceError{
				Name: "a1b2c3d4e5f6",
			},
			want: "git reference in module source path unsupported a1b2c3d4e5f6",
		},
		{
			name: "tag reference",
			m: &gitReferenceError{
				Name: "refs/tags/v1.0.0",
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
				Name: "",
			},
			want: "repository not formatted correctly ",
		},
		{
			name: "invalid path format",
			m: &repositoryFormatError{
				Name: "invalid/repo/path///",
			},
			want: "repository not formatted correctly invalid/repo/path///",
		},
		{
			name: "special characters",
			m: &repositoryFormatError{
				Name: "@#$%^&*",
			},
			want: "repository not formatted correctly @#$%^&*",
		},
		{
			name: "url format",
			m: &repositoryFormatError{
				Name: "https://github.com/user/repo.git",
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
				Name: "   ",
			},
			want: "repository not formatted correctly    ",
		},
		{
			name: "very long path",
			m: &repositoryFormatError{
				Name: "org/extremely/long/repository/path/that/exceeds/normal/length/limits/and/tests/boundary/conditions",
			},
			want: "repository not formatted correctly org/extremely/long/repository/path/that/exceeds/normal/length/limits/and/tests/boundary/conditions",
		},
		{
			name: "unicode characters",
			m: &repositoryFormatError{
				Name: "организация/репозиторий",
			},
			want: "repository not formatted correctly организация/репозиторий",
		},
		{
			name: "numeric repository",
			m: &repositoryFormatError{
				Name: "12345/67890",
			},
			want: "repository not formatted correctly 12345/67890",
		},
		{
			name: "mixed case with symbols",
			m: &repositoryFormatError{
				Name: "My-Org/Repo_Name-123",
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
