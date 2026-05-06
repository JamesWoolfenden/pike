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
	_ = os.Setenv("GCP_PROJECT", "pike-477416")
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
			"resource\"google_project_iam_custom_role\"\"terraform_pike\"{project=\"pike-477416\"role_id=\"terraform_pike\"title=\"terraform_pike\"description=\"Auserwithleastprivileges\"permissions=[\"bigquery.datasets.create\",\"bigquery.jobs.create\"]}",
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

			got, err := GCPPolicy(tt.args.permissions, "", "")

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

// allProjectEnvVars is the full list of env vars getCurrentProject checks.
var allProjectEnvVars = []string{
	"GOOGLE_CLOUD_PROJECT",
	"CLOUDSDK_CORE_PROJECT",
	"GOOGLE_PROJECT",
	"GCLOUD_PROJECT",
	"GCP_PROJECT",
}

func saveAllProjectEnvVars() map[string]string {
	saved := make(map[string]string, len(allProjectEnvVars))
	for _, v := range allProjectEnvVars {
		saved[v] = os.Getenv(v)
	}
	return saved
}

func restoreAllProjectEnvVars(saved map[string]string) {
	for k, v := range saved {
		_ = os.Setenv(k, v)
	}
}

func clearAllProjectEnvVars() {
	for _, v := range allProjectEnvVars {
		_ = os.Unsetenv(v)
	}
}

func TestGetCurrentProject_EnvironmentVariables(t *testing.T) {
	saved := saveAllProjectEnvVars()
	defer restoreAllProjectEnvVars(saved)

	tests := []struct {
		name     string
		env      map[string]string
		expected string
	}{
		{
			name: "GOOGLE_CLOUD_PROJECT takes precedence",
			env: map[string]string{
				"GOOGLE_CLOUD_PROJECT":  "test-project-1",
				"CLOUDSDK_CORE_PROJECT": "test-project-x",
				"GOOGLE_PROJECT":        "test-project-2",
				"GCP_PROJECT":           "test-project-3",
			},
			expected: "test-project-1",
		},
		{
			name: "CLOUDSDK_CORE_PROJECT second priority",
			env: map[string]string{
				"CLOUDSDK_CORE_PROJECT": "test-project-x",
				"GOOGLE_PROJECT":        "test-project-2",
				"GCP_PROJECT":           "test-project-3",
			},
			expected: "test-project-x",
		},
		{
			name: "GOOGLE_PROJECT third priority",
			env: map[string]string{
				"GOOGLE_PROJECT": "test-project-2",
				"GCP_PROJECT":    "test-project-3",
			},
			expected: "test-project-2",
		},
		{
			name: "GCLOUD_PROJECT fourth priority",
			env: map[string]string{
				"GCLOUD_PROJECT": "test-project-y",
				"GCP_PROJECT":    "test-project-3",
			},
			expected: "test-project-y",
		},
		{
			name:     "GCP_PROJECT when others empty",
			env:      map[string]string{"GCP_PROJECT": "test-project-3"},
			expected: "test-project-3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clearAllProjectEnvVars()
			for k, v := range tt.env {
				_ = os.Setenv(k, v)
			}

			project, err := getCurrentProject("")
			if err != nil {
				t.Fatalf("Expected no error, got: %v", err)
			}

			if project != tt.expected {
				t.Errorf("Expected project %s, got %s", tt.expected, project)
			}
		})
	}
}

func TestGetCurrentProject_GcloudConfigFile(t *testing.T) {
	saved := saveAllProjectEnvVars()
	originalHome := os.Getenv("HOME")
	originalAppData := os.Getenv("APPDATA")
	originalGoogleAppCreds := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")

	defer func() {
		restoreAllProjectEnvVars(saved)
		_ = os.Setenv("HOME", originalHome)
		_ = os.Setenv("APPDATA", originalAppData)
		_ = os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", originalGoogleAppCreds)
	}()

	// Clear all env vars that could short-circuit the config file path.
	// GOOGLE_APPLICATION_CREDENTIALS points to an absolute key file and
	// is unaffected by HOME/APPDATA redirection, so it must be cleared too.
	clearAllProjectEnvVars()
	_ = os.Unsetenv("GOOGLE_APPLICATION_CREDENTIALS")

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
	err := os.MkdirAll(filepath.Dir(configPath), 0o755)
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

	err = os.WriteFile(configPath, []byte(configContent), 0o644)
	if err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	project, err := getCurrentProject("")
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if project != "test-config-project" {
		t.Errorf("Expected project 'test-config-project', got '%s'", project)
	}
}

func TestGetCurrentProject_EmptyEnvironmentVariables(t *testing.T) {
	saved := saveAllProjectEnvVars()
	defer restoreAllProjectEnvVars(saved)

	// Test empty string environment variables (should be treated as not set)
	for _, v := range allProjectEnvVars {
		_ = os.Setenv(v, "")
	}

	// This will likely fail due to no credentials or config file, but we're testing the logic
	_, err := getCurrentProject("")
	// We expect an error since no valid project source is available
	if err == nil {
		t.Log("No error returned - likely found valid credentials or config file")
	}
}

func TestGetCurrentProject_MissingConfigFile(t *testing.T) {
	saved := saveAllProjectEnvVars()
	originalHome := os.Getenv("HOME")
	originalAppData := os.Getenv("APPDATA")
	originalGoogleAppCreds := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")

	defer func() {
		restoreAllProjectEnvVars(saved)
		_ = os.Setenv("HOME", originalHome)
		_ = os.Setenv("APPDATA", originalAppData)
		_ = os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", originalGoogleAppCreds)
	}()

	clearAllProjectEnvVars()
	_ = os.Unsetenv("GOOGLE_APPLICATION_CREDENTIALS")

	// Set HOME/APPDATA to non-existent directory
	tempDir := t.TempDir()
	nonExistentDir := filepath.Join(tempDir, "nonexistent")

	if runtime.GOOS != "windows" {
		_ = os.Setenv("HOME", nonExistentDir)
	} else {
		_ = os.Setenv("APPDATA", nonExistentDir)
	}

	_, err := getCurrentProject("")
	if err == nil {
		t.Log("No error returned - likely found valid credentials")
	}
}

func TestGetCurrentProject_InvalidConfigFile(t *testing.T) {
	saved := saveAllProjectEnvVars()
	originalHome := os.Getenv("HOME")
	originalAppData := os.Getenv("APPDATA")
	originalGoogleAppCreds := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")

	defer func() {
		restoreAllProjectEnvVars(saved)
		_ = os.Setenv("HOME", originalHome)
		_ = os.Setenv("APPDATA", originalAppData)
		_ = os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", originalGoogleAppCreds)
	}()

	clearAllProjectEnvVars()
	_ = os.Unsetenv("GOOGLE_APPLICATION_CREDENTIALS")

	tempDir := t.TempDir()

	var configPath string
	if runtime.GOOS != "windows" {
		_ = os.Setenv("HOME", tempDir)
		configPath = filepath.Join(tempDir, ".config", "gcloud", "configurations", "config_default")
	} else {
		_ = os.Setenv("APPDATA", tempDir)
		configPath = filepath.Join(tempDir, "gcloud", "configurations", "config_default")
	}

	err := os.MkdirAll(filepath.Dir(configPath), 0o755)
	if err != nil {
		t.Fatalf("Failed to create config directory: %v", err)
	}

	invalidConfigContent := `[core
project = test-project
invalid ini format
`

	err = os.WriteFile(configPath, []byte(invalidConfigContent), 0o644)
	if err != nil {
		t.Fatalf("Failed to write invalid config file: %v", err)
	}

	// Invalid ini falls through to credentials; no credentials in CI so expect error
	_, err = getCurrentProject("")
	if err == nil {
		t.Log("No error returned - ini parser was lenient or valid credentials found")
	}
}

func TestGetCurrentProject_ConfigFileWithoutProject(t *testing.T) {
	saved := saveAllProjectEnvVars()
	originalHome := os.Getenv("HOME")
	originalAppData := os.Getenv("APPDATA")
	originalGoogleAppCreds := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")

	defer func() {
		restoreAllProjectEnvVars(saved)
		_ = os.Setenv("HOME", originalHome)
		_ = os.Setenv("APPDATA", originalAppData)
		_ = os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", originalGoogleAppCreds)
	}()

	clearAllProjectEnvVars()
	_ = os.Unsetenv("GOOGLE_APPLICATION_CREDENTIALS")

	tempDir := t.TempDir()

	var configPath string
	if runtime.GOOS != "windows" {
		_ = os.Setenv("HOME", tempDir)
		configPath = filepath.Join(tempDir, ".config", "gcloud", "configurations", "config_default")
	} else {
		_ = os.Setenv("APPDATA", tempDir)
		configPath = filepath.Join(tempDir, "gcloud", "configurations", "config_default")
	}

	err := os.MkdirAll(filepath.Dir(configPath), 0o755)
	if err != nil {
		t.Fatalf("Failed to create config directory: %v", err)
	}

	configContent := `[core]
account = test@example.com

[compute]
region = us-central1
zone = us-central1-a
`

	err = os.WriteFile(configPath, []byte(configContent), 0o644)
	if err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	// Config exists but has no project; falls through to credentials which are absent → error
	_, err = getCurrentProject("")
	if err == nil {
		t.Log("No error returned - likely found valid credentials")
	}
}

func TestGetProjectFromTerraformProvider(t *testing.T) {
	t.Parallel()

	tempDir := t.TempDir()

	tfContent := `provider "google" {
  project = "my-gcp-project-123"
  region  = "us-central1"
}

resource "google_storage_bucket" "example" {
  name = "example"
}
`

	err := os.WriteFile(filepath.Join(tempDir, "main.tf"), []byte(tfContent), 0o644)
	if err != nil {
		t.Fatalf("Failed to write tf file: %v", err)
	}

	project := getProjectFromTerraformProvider(tempDir)
	if project != "my-gcp-project-123" {
		t.Errorf("Expected 'my-gcp-project-123', got '%s'", project)
	}
}

func TestGetProjectFromTerraformProvider_NoProvider(t *testing.T) {
	t.Parallel()

	tempDir := t.TempDir()

	tfContent := `resource "google_storage_bucket" "example" {
  name = "example"
}
`

	err := os.WriteFile(filepath.Join(tempDir, "main.tf"), []byte(tfContent), 0o644)
	if err != nil {
		t.Fatalf("Failed to write tf file: %v", err)
	}

	project := getProjectFromTerraformProvider(tempDir)
	if project != "" {
		t.Errorf("Expected empty string, got '%s'", project)
	}
}

func TestGetProjectFromTerraformProvider_EmptyDir(t *testing.T) {
	t.Parallel()

	project := getProjectFromTerraformProvider("")
	if project != "" {
		t.Errorf("Expected empty string for empty dirName, got '%s'", project)
	}
}
