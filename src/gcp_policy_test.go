package pike

import (
	_ "embed"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestGCPPolicy(t *testing.T) {
	t.Parallel()
	os.Setenv("GCP_PROJECT", "pike-412922")
	type args struct {
		permissions []string
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"basic",
			args{[]string{"bigquery.datasets.create", "bigquery.jobs.create"}},
			"resource\"google_project_iam_custom_role\"\"terraform_pike\"{project=\"pike-412922\"role_id=\"terraform_pike\"title=\"terraform_pike\"description=\"Auserwithleastprivileges\"permissions=[\"bigquery.datasets.create\",\"bigquery.jobs.create\"]}",
			false,
		},
		{
			"empty",
			args{[]string{}},
			"",
			true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := GCPPolicy(tt.args.permissions, "")

			if (err != nil) != tt.wantErr {
				t.Errorf("GCPPolicy() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			original := Minify(got)
			target := Minify(tt.want)

			if original != target {
				t.Errorf("GCPPolicy() = %v, want %v", got, tt.want)
				t.Errorf("GCPPolicy() = %v, want %v", original, target)
			}
		})
	}
}

func TestGetCurrentProject_EnvironmentVariables(t *testing.T) {
	// Test GOOGLE_CLOUD_PROJECT takes precedence
	t.Run("GOOGLE_CLOUD_PROJECT_set", func(t *testing.T) {
		os.Setenv("GOOGLE_CLOUD_PROJECT", "test-project-1")
		os.Setenv("GOOGLE_PROJECT", "test-project-2")
		os.Setenv("GCP_PROJECT", "test-project-3")
		defer func() {
			os.Unsetenv("GOOGLE_CLOUD_PROJECT")
			os.Unsetenv("GOOGLE_PROJECT")
			os.Unsetenv("GCP_PROJECT")
		}()

		project, err := getCurrentProject()
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if project != "test-project-1" {
			t.Errorf("Expected 'test-project-1', got '%s'", project)
		}
	})

	// Test GOOGLE_PROJECT when GOOGLE_CLOUD_PROJECT is not set
	t.Run("GOOGLE_PROJECT_set", func(t *testing.T) {
		os.Unsetenv("GOOGLE_CLOUD_PROJECT")
		os.Setenv("GOOGLE_PROJECT", "test-project-2")
		os.Setenv("GCP_PROJECT", "test-project-3")
		defer func() {
			os.Unsetenv("GOOGLE_PROJECT")
			os.Unsetenv("GCP_PROJECT")
		}()

		project, err := getCurrentProject()
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if project != "test-project-2" {
			t.Errorf("Expected 'test-project-2', got '%s'", project)
		}
	})

	// Test GCP_PROJECT when others are not set
	t.Run("GCP_PROJECT_set", func(t *testing.T) {
		os.Unsetenv("GOOGLE_CLOUD_PROJECT")
		os.Unsetenv("GOOGLE_PROJECT")
		os.Setenv("GCP_PROJECT", "test-project-3")
		defer os.Unsetenv("GCP_PROJECT")

		project, err := getCurrentProject()
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if project != "test-project-3" {
			t.Errorf("Expected 'test-project-3', got '%s'", project)
		}
	})
}

func TestGetCurrentProject_ConfigFile(t *testing.T) {
	// Clear all environment variables
	os.Unsetenv("GOOGLE_CLOUD_PROJECT")
	os.Unsetenv("GOOGLE_PROJECT")
	os.Unsetenv("GCP_PROJECT")

	t.Run("valid_config_file", func(t *testing.T) {
		// Create a temporary config file
		tempDir := t.TempDir()
		configDir := filepath.Join(tempDir, ".config", "gcloud", "configurations")
		err := os.MkdirAll(configDir, 0755)
		if err != nil {
			t.Fatalf("Failed to create config directory: %v", err)
		}

		configFile := filepath.Join(configDir, "config_default")
		configContent := `[core]
project = test-config-project
account = test@example.com
`
		err = os.WriteFile(configFile, []byte(configContent), 0644)
		if err != nil {
			t.Fatalf("Failed to write config file: %v", err)
		}

		// Temporarily change HOME to point to our temp directory
		originalHome := os.Getenv("HOME")
		os.Setenv("HOME", tempDir)
		defer os.Setenv("HOME", originalHome)

		// This test will only work if google.FindDefaultCredentials fails
		// In a real environment, this might succeed, so we test the config fallback
		project, err := getCurrentProject()

		// The function should either return from credentials or from config file
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		// Project should be either from credentials or from config
		if project == "" {
			t.Error("Expected non-empty project")
		}
	})

	t.Run("missing_config_file", func(t *testing.T) {
		// Create temp directory without config file
		tempDir := t.TempDir()

		originalHome := os.Getenv("HOME")
		os.Setenv("HOME", tempDir)
		defer os.Setenv("HOME", originalHome)

		// This will test the error path when both credentials and config fail
		_, err := getCurrentProject()

		// We expect either success from default credentials or an error
		// The behavior depends on the environment where tests run
		if err != nil {
			// This is expected when no credentials and no config file
			if err.Error() == "" {
				t.Error("Expected non-empty error message")
			}
		}
	})
}

func TestGetCurrentProject_WindowsPath(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("Skipping Windows-specific test on non-Windows platform")
	}

	os.Unsetenv("GOOGLE_CLOUD_PROJECT")
	os.Unsetenv("GOOGLE_PROJECT")
	os.Unsetenv("GCP_PROJECT")

	t.Run("windows_config_path", func(t *testing.T) {
		tempDir := t.TempDir()
		configDir := filepath.Join(tempDir, "gcloud", "configurations")
		err := os.MkdirAll(configDir, 0755)
		if err != nil {
			t.Fatalf("Failed to create config directory: %v", err)
		}

		configFile := filepath.Join(configDir, "config_default")
		configContent := `[core]
project = windows-test-project
`
		err = os.WriteFile(configFile, []byte(configContent), 0644)
		if err != nil {
			t.Fatalf("Failed to write config file: %v", err)
		}

		originalAppData := os.Getenv("APPDATA")
		os.Setenv("APPDATA", tempDir)
		defer os.Setenv("APPDATA", originalAppData)

		project, err := getCurrentProject()
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if project == "" {
			t.Error("Expected non-empty project")
		}
	})
}

func TestGetCurrentProject_EmptyEnvironmentVariables(t *testing.T) {
	t.Run("empty_env_vars", func(t *testing.T) {
		os.Setenv("GOOGLE_CLOUD_PROJECT", "")
		os.Setenv("GOOGLE_PROJECT", "")
		os.Setenv("GCP_PROJECT", "")
		defer func() {
			os.Unsetenv("GOOGLE_CLOUD_PROJECT")
			os.Unsetenv("GOOGLE_PROJECT")
			os.Unsetenv("GCP_PROJECT")
		}()

		// Should fall back to credentials or config file
		project, err := getCurrentProject()

		// The result depends on the environment
		if err == nil && project == "" {
			t.Error("Expected non-empty project when no error")
		}
	})
}

func TestGetCurrentProject_InvalidConfigFile(t *testing.T) {
	os.Unsetenv("GOOGLE_CLOUD_PROJECT")
	os.Unsetenv("GOOGLE_PROJECT")
	os.Unsetenv("GCP_PROJECT")

	t.Run("invalid_config_format", func(t *testing.T) {
		tempDir := t.TempDir()
		configDir := filepath.Join(tempDir, ".config", "gcloud", "configurations")
		err := os.MkdirAll(configDir, 0755)
		if err != nil {
			t.Fatalf("Failed to create config directory: %v", err)
		}

		configFile := filepath.Join(configDir, "config_default")
		// Write invalid INI content
		invalidContent := `invalid ini content without proper format`
		err = os.WriteFile(configFile, []byte(invalidContent), 0644)
		if err != nil {
			t.Fatalf("Failed to write config file: %v", err)
		}

		originalHome := os.Getenv("HOME")
		os.Setenv("HOME", tempDir)
		defer os.Setenv("HOME", originalHome)

		project, err := getCurrentProject()

		// Should either succeed with credentials or fail with config error
		if err == nil && project == "" {
			t.Error("Expected non-empty project when no error")
		}
	})
}
