package pike

import (
	"os"
	"path/filepath"
	"slices"
	"testing"
)

func TestGCPQuotaPermissions(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name, configuration, filter string
		want                        bool
	}{
		{"enabled", `provider "google" {
  billing_project = "quota-project"
  user_project_override = true
}`, "", true},
		{"variable billing project", `provider "google" {
  billing_project = var.quota_project
  user_project_override = true
}`, "google", true},
		{"resource project fallback", `provider "google" { user_project_override = true }`, "gcp", true},
		{"beta alias", `provider "google-beta" {
  alias = "quota"
  user_project_override = true
}`, "google", true},
		{"disabled", `provider "google" {
  billing_project = "quota-project"
  user_project_override = false
}`, "", false},
		{"omitted", `provider "google" { billing_project = "quota-project" }`, "", false},
		{"unknown", `provider "google" { user_project_override = var.override }`, "", true},
		{"constant expression", `provider "google" { user_project_override = !false }`, "", true},
		{"string true", `provider "google" { user_project_override = "true" }`, "", true},
		{"string false", `provider "google" { user_project_override = "false" }`, "", false},
		{"null", `provider "google" { user_project_override = null }`, "", false},
		{"invalid type", `provider "google" { user_project_override = [] }`, "", false},
		{"unrelated provider", `provider "aws" { user_project_override = true }`, "", false},
		{"aws filter", `provider "google" { user_project_override = true }`, "aws", false},
		{"azure filter", `provider "google" { user_project_override = true }`, "azure", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			dir := t.TempDir()
			writeGCPFixture(t, dir, "provider.tf", tt.configuration)
			writeGCPFixture(t, dir, "main.tf", `resource "google_storage_bucket" "example" {
  name = "example"
}`)
			bag, err := MakePermissionBag(dir, nil, false, tt.filter, false)
			if err != nil {
				t.Fatal(err)
			}
			for phase, permissions := range map[string][]string{"apply": bag.GCP, "plan": bag.PlanGCP} {
				if got := slices.Contains(permissions, "serviceusage.services.use"); got != tt.want {
					t.Errorf("%s quota permission = %v, want %v: %v", phase, got, tt.want, permissions)
				}
			}
		})
	}
}

func TestGCPQuotaPermissionsSingleFileAndLocalModule(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	moduleDir := filepath.Join(dir, "child")
	if err := os.Mkdir(moduleDir, 0o750); err != nil {
		t.Fatal(err)
	}
	child := writeGCPFixture(t, moduleDir, "main.tf", `provider "google" {
  user_project_override = true
}
resource "google_storage_bucket" "example" { name = "example" }`)
	writeGCPFixture(t, dir, "main.tf", `module "child" { source = "./child" }`)

	for _, file := range []*string{nil, &child} {
		bag, err := MakePermissionBag(dir, file, false, "gcp", false)
		if err != nil {
			t.Fatal(err)
		}
		if !slices.Contains(bag.GCP, "serviceusage.services.use") ||
			!slices.Contains(bag.PlanGCP, "serviceusage.services.use") {
			t.Errorf("missing quota permission: %+v", bag)
		}
	}
}

func writeGCPFixture(t *testing.T, dir, name, content string) string {
	t.Helper()
	path := filepath.Join(dir, name)
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}
	return path
}
