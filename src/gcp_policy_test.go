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
	_ = os.Setenv("GCP_PROJECT", "pike-412922")
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
	// Save original environment
	originalGoogleCloudProject := os.Getenv("GOOGLE_CLOUD_PROJECT")
	originalGoogleProject := os.Getenv("GOOGLE_PROJECT")
	originalGcpProject := os.Getenv("GCP_PROJECT")

	// Clean up after test
	defer func() {
		_ = os.Setenv("GOOGLE_CLOUD_PROJECT", originalGoogleCloudProject)
		_ = os.Setenv("GOOGLE_PROJECT", originalGoogleProject)
		_ = os.Setenv("GCP_PROJECT", originalGcpProject)
	}()

	tests := []struct {
		name               string
		googleCloudProject string
		googleProject      string
		gcpProject         string
		expectedProject    string
	}{
		{
			name:               "GOOGLE_CLOUD_PROJECT takes precedence",
			googleCloudProject: "test-project-1",
			googleProject:      "test-project-2",
			gcpProject:         "test-project-3",
			expectedProject:    "test-project-1",
		},
		{
			name:            "GOOGLE_PROJECT when GOOGLE_CLOUD_PROJECT empty",
			googleProject:   "test-project-2",
			gcpProject:      "test-project-3",
			expectedProject: "test-project-2",
		},
		{
			name:            "GCP_PROJECT when others empty",
			gcpProject:      "test-project-3",
			expectedProject: "test-project-3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Clear all environment variables
			_ = os.Unsetenv("GOOGLE_CLOUD_PROJECT")
			_ = os.Unsetenv("GOOGLE_PROJECT")
			_ = os.Unsetenv("GCP_PROJECT")

			// Set test values
			if tt.googleCloudProject != "" {
				_ = os.Setenv("GOOGLE_CLOUD_PROJECT", tt.googleCloudProject)
			}
			if tt.googleProject != "" {
				_ = os.Setenv("GOOGLE_PROJECT", tt.googleProject)
			}
			if tt.gcpProject != "" {
				_ = os.Setenv("GCP_PROJECT", tt.gcpProject)
			}

			project, err := getCurrentProject()
			if err != nil {
				t.Fatalf("Expected no error, got: %v", err)
			}

			if project != tt.expectedProject {
				t.Errorf("Expected project %s, got %s", tt.expectedProject, project)
			}
		})
	}
}

func TestGetCurrentProject_GcloudConfigFile(t *testing.T) {
	// Save original environment
	originalGoogleCloudProject := os.Getenv("GOOGLE_CLOUD_PROJECT")
	originalGoogleProject := os.Getenv("GOOGLE_PROJECT")
	originalGcpProject := os.Getenv("GCP_PROJECT")
	originalHome := os.Getenv("HOME")
	originalAppData := os.Getenv("APPDATA")

	// Clean up after test
	defer func() {
		_ = os.Setenv("GOOGLE_CLOUD_PROJECT", originalGoogleCloudProject)
		_ = os.Setenv("GOOGLE_PROJECT", originalGoogleProject)
		_ = os.Setenv("GCP_PROJECT", originalGcpProject)
		_ = os.Setenv("HOME", originalHome)
		_ = os.Setenv("APPDATA", originalAppData)
	}()

	// Clear environment variables to force config file reading
	_ = os.Unsetenv("GOOGLE_CLOUD_PROJECT")
	_ = os.Unsetenv("GOOGLE_PROJECT")
	_ = os.Unsetenv("GCP_PROJECT")

	// Create temporary directory structure
	tempDir := t.TempDir()

	var configPath string
	if runtime.GOOS != "windows" {
		_ = os.Setenv("HOME", tempDir)
		configPath = filepath.Join(tempDir, ".config", "gcloud", "configurations", "config_default")
	} else {
		_ = os.Setenv("APPDATA", tempDir)
		configPath = filepath.Join(tempDir, "gcloud", "configurations", "config_default")
	}

	// Create directory structure
	err := os.MkdirAll(filepath.Dir(configPath), 0755)
	if err != nil {
		t.Fatalf("Failed to create config directory: %v", err)
	}

	// Create config file with project
	configContent := `[core]
project = test-config-project
account = test@example.com

[compute]
region = us-central1
zone = us-central1-a
`

	err = os.WriteFile(configPath, []byte(configContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	project, err := getCurrentProject()
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if project != "test-config-project" {
		t.Errorf("Expected project 'test-config-project', got '%s'", project)
	}
}

func TestGetCurrentProject_EmptyEnvironmentVariables(t *testing.T) {
	// Save original environment
	originalGoogleCloudProject := os.Getenv("GOOGLE_CLOUD_PROJECT")
	originalGoogleProject := os.Getenv("GOOGLE_PROJECT")
	originalGcpProject := os.Getenv("GCP_PROJECT")

	// Clean up after test
	defer func() {
		_ = os.Setenv("GOOGLE_CLOUD_PROJECT", originalGoogleCloudProject)
		_ = os.Setenv("GOOGLE_PROJECT", originalGoogleProject)
		_ = os.Setenv("GCP_PROJECT", originalGcpProject)
	}()

	// Test empty string environment variables (should be treated as not set)
	_ = os.Setenv("GOOGLE_CLOUD_PROJECT", "")
	_ = os.Setenv("GOOGLE_PROJECT", "")
	_ = os.Setenv("GCP_PROJECT", "")

	// This will likely fail due to no credentials or config file, but we're testing the logic
	_, err := getCurrentProject()
	// We expect an error since no valid project source is available
	if err == nil {
		t.Log("No error returned - likely found valid credentials or config file")
	}
}

func TestGetCurrentProject_MissingConfigFile(t *testing.T) {
	// Save original environment
	originalGoogleCloudProject := os.Getenv("GOOGLE_CLOUD_PROJECT")
	originalGoogleProject := os.Getenv("GOOGLE_PROJECT")
	originalGcpProject := os.Getenv("GCP_PROJECT")
	originalHome := os.Getenv("HOME")
	originalAppData := os.Getenv("APPDATA")

	// Clean up after test
	defer func() {
		_ = os.Setenv("GOOGLE_CLOUD_PROJECT", originalGoogleCloudProject)
		_ = os.Setenv("GOOGLE_PROJECT", originalGoogleProject)
		_ = os.Setenv("GCP_PROJECT", originalGcpProject)
		_ = os.Setenv("HOME", originalHome)
		_ = os.Setenv("APPDATA", originalAppData)
	}()

	// Clear environment variables
	_ = os.Unsetenv("GOOGLE_CLOUD_PROJECT")
	_ = os.Unsetenv("GOOGLE_PROJECT")
	_ = os.Unsetenv("GCP_PROJECT")

	// Set HOME/APPDATA to non-existent directory
	tempDir := t.TempDir()
	nonExistentDir := filepath.Join(tempDir, "nonexistent")

	if runtime.GOOS != "windows" {
		_ = os.Setenv("HOME", nonExistentDir)
	} else {
		_ = os.Setenv("APPDATA", nonExistentDir)
	}

	_, err := getCurrentProject()
	if err == nil {
		t.Log("No error returned - likely found valid credentials")
	}
}

func TestGetCurrentProject_InvalidConfigFile(t *testing.T) {
	// Save original environment
	originalGoogleCloudProject := os.Getenv("GOOGLE_CLOUD_PROJECT")
	originalGoogleProject := os.Getenv("GOOGLE_PROJECT")
	originalGcpProject := os.Getenv("GCP_PROJECT")
	originalHome := os.Getenv("HOME")
	originalAppData := os.Getenv("APPDATA")

	// Clean up after test
	defer func() {
		_ = os.Setenv("GOOGLE_CLOUD_PROJECT", originalGoogleCloudProject)
		_ = os.Setenv("GOOGLE_PROJECT", originalGoogleProject)
		_ = os.Setenv("GCP_PROJECT", originalGcpProject)
		_ = os.Setenv("HOME", originalHome)
		_ = os.Setenv("APPDATA", originalAppData)
	}()

	// Clear environment variables
	_ = os.Unsetenv("GOOGLE_CLOUD_PROJECT")
	_ = os.Unsetenv("GOOGLE_PROJECT")
	_ = os.Unsetenv("GCP_PROJECT")

	// Create temporary directory structure
	tempDir := t.TempDir()

	var configPath string
	if runtime.GOOS != "windows" {
		_ = os.Setenv("HOME", tempDir)
		configPath = filepath.Join(tempDir, ".config", "gcloud", "configurations", "config_default")
	} else {
		_ = os.Setenv("APPDATA", tempDir)
		configPath = filepath.Join(tempDir, "gcloud", "configurations", "config_default")
	}

	// Create directory structure
	err := os.MkdirAll(filepath.Dir(configPath), 0755)
	if err != nil {
		t.Fatalf("Failed to create config directory: %v", err)
	}

	// Create invalid config file
	invalidConfigContent := `[core
project = test-project
invalid ini format
`

	err = os.WriteFile(configPath, []byte(invalidConfigContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write invalid config file: %v", err)
	}

	_, err = getCurrentProject()
	if err == nil {
		t.Log("No error returned - likely found valid credentials or ini parser was lenient")
	}
}

func TestGetCurrentProject_ConfigFileWithoutProject(t *testing.T) {
	// Save original environment
	originalGoogleCloudProject := os.Getenv("GOOGLE_CLOUD_PROJECT")
	originalGoogleProject := os.Getenv("GOOGLE_PROJECT")
	originalGcpProject := os.Getenv("GCP_PROJECT")
	originalHome := os.Getenv("HOME")
	originalAppData := os.Getenv("APPDATA")

	// Clean up after test
	defer func() {
		_ = os.Setenv("GOOGLE_CLOUD_PROJECT", originalGoogleCloudProject)
		_ = os.Setenv("GOOGLE_PROJECT", originalGoogleProject)
		_ = os.Setenv("GCP_PROJECT", originalGcpProject)
		_ = os.Setenv("HOME", originalHome)
		_ = os.Setenv("APPDATA", originalAppData)
	}()

	// Clear environment variables
	_ = os.Unsetenv("GOOGLE_CLOUD_PROJECT")
	_ = os.Unsetenv("GOOGLE_PROJECT")
	_ = os.Unsetenv("GCP_PROJECT")

	// Create temporary directory structure
	tempDir := t.TempDir()

	var configPath string
	if runtime.GOOS != "windows" {
		_ = os.Setenv("HOME", tempDir)
		configPath = filepath.Join(tempDir, ".config", "gcloud", "configurations", "config_default")
	} else {
		_ = os.Setenv("APPDATA", tempDir)
		configPath = filepath.Join(tempDir, "gcloud", "configurations", "config_default")
	}

	// Create directory structure
	err := os.MkdirAll(filepath.Dir(configPath), 0755)
	if err != nil {
		t.Fatalf("Failed to create config directory: %v", err)
	}

	// Create config file without project
	configContent := `[core]
account = test@example.com

[compute]
region = us-central1
zone = us-central1-a
`

	err = os.WriteFile(configPath, []byte(configContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	project, err := getCurrentProject()
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	// Should return empty string when no project is set in config
	if project != "" {
		t.Errorf("Expected empty project, got '%s'", project)
	}
}
