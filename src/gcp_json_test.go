package pike

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestGCPJSONOutput(t *testing.T) {
	t.Setenv("GOOGLE_CLOUD_PROJECT", "test-project")
	policy, err := GetPolicy(Sorted{
		GCP:     []string{"storage.buckets.update", "storage.buckets.get", "storage.buckets.get"},
		PlanGCP: []string{"storage.buckets.get"},
	}, false, "", "")
	if err != nil {
		t.Fatal(err)
	}
	assertGCPJSONPermissions(t, policy.AsString("json"), "", []string{"storage.buckets.get", "storage.buckets.update"})
	assertGCPJSONPermissions(t, policy.AsTwoRoleString("json"), "apply", []string{"storage.buckets.get", "storage.buckets.update"})
	assertGCPJSONPermissions(t, policy.AsTwoRoleString("JSON"), "plan", []string{"storage.buckets.get"})
	if !strings.Contains(policy.AsTwoRoleString("terraform"), `resource "google_project_iam_custom_role"`) {
		t.Fatal("Terraform output lost the GCP roles")
	}
}

func TestAWSJSONCompatibility(t *testing.T) {
	t.Parallel()
	policy, err := GetPolicy(Sorted{
		AWS:     []string{"s3:GetObject"},
		PlanAWS: []string{"s3:GetObject"},
	}, false, "", "")
	if err != nil {
		t.Fatal(err)
	}
	if policy.AsString("json") != policy.AWS.JSONOut+"\n" {
		t.Fatal("AWS legacy JSON changed")
	}
	var roles map[string]json.RawMessage
	if err := json.Unmarshal([]byte(policy.AsTwoRoleString("json")), &roles); err != nil {
		t.Fatal(err)
	}
	if len(roles) != 2 {
		t.Fatalf("AWS two-role JSON shape changed: %v", roles)
	}
	for _, phase := range []string{"apply", "plan"} {
		var document Policy
		if err := json.Unmarshal(roles[phase], &document); err != nil {
			t.Fatal(err)
		}
		if document.Version != "2012-10-17" || len(document.Statements) != 1 {
			t.Errorf("AWS %s document changed: %+v", phase, document)
		}
	}
}

func TestGCPJSONWithoutPlanPermissions(t *testing.T) {
	t.Setenv("GOOGLE_CLOUD_PROJECT", "test-project")
	policy, err := GetPolicy(Sorted{GCP: []string{"storage.buckets.get"}}, false, "", "")
	if err != nil {
		t.Fatal(err)
	}
	output := policy.AsTwoRoleString("json")
	assertGCPJSONPermissions(t, output, "apply", []string{"storage.buckets.get"})
	var roles map[string]json.RawMessage
	if err := json.Unmarshal([]byte(output), &roles); err != nil {
		t.Fatal(err)
	}
	if _, exists := roles["plan"]; exists {
		t.Fatal("missing plan permissions must not become an invented plan role")
	}
}

func TestGCPJSONMixedProviders(t *testing.T) {
	t.Setenv("GOOGLE_CLOUD_PROJECT", "test-project")
	policy, err := GetPolicy(Sorted{
		AWS:     []string{"s3:GetObject"},
		PlanAWS: []string{"s3:GetObject"},
		GCP:     []string{"storage.buckets.get"},
		PlanGCP: []string{"storage.buckets.get"},
	}, false, "", "")
	if err != nil {
		t.Fatal(err)
	}
	for _, legacy := range []bool{false, true} {
		output := policy.AsTwoRoleString("json")
		if legacy {
			output = policy.AsString("json")
		}
		var providers map[string]json.RawMessage
		if err := json.Unmarshal([]byte(output), &providers); err != nil {
			t.Fatalf("invalid JSON: %v\n%s", err, output)
		}
		if len(providers["aws"]) == 0 || len(providers["gcp"]) == 0 {
			t.Fatalf("mixed output must preserve both providers: %s", output)
		}
		phase := "plan"
		if legacy {
			phase = ""
		}
		assertGCPJSONPermissions(t, string(providers["gcp"]), phase, []string{"storage.buckets.get"})
	}
}

func TestGCPJSONScanWrite(t *testing.T) {
	t.Setenv("GOOGLE_CLOUD_PROJECT", "test-project")
	for _, legacy := range []bool{false, true} {
		for _, explicit := range []bool{false, true} {
			dir := t.TempDir()
			writeGCPJSONFixture(t, dir, "main.tf", `resource "google_storage_bucket" "example" {
  name = "example"
}`)
			file := ""
			wantPath := filepath.Join(dir, ".pike", "pike.two-role.json")
			phase := "plan"
			if legacy {
				wantPath = filepath.Join(dir, ".pike", "pike.generated_policy.json")
				phase = ""
			}
			if explicit {
				file = filepath.Join(dir, "permissions.json")
				wantPath = file
			}
			if err := Scan(dir, "json", nil, false, true, false, "gcp", file, "", false, legacy); err != nil {
				t.Fatal(err)
			}
			content, err := os.ReadFile(wantPath)
			if err != nil {
				t.Fatal(err)
			}
			assertGCPJSONPermissions(t, string(content), phase, nil)
			if _, err := os.Stat(filepath.Join(dir, ".pike", "pike.two-role.tf")); !os.IsNotExist(err) {
				t.Fatal("JSON scan unexpectedly wrote Terraform")
			}
		}
	}
}

func writeGCPJSONFixture(t *testing.T, dir, name, content string) {
	t.Helper()
	if err := os.WriteFile(filepath.Join(dir, name), []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}
}

func assertGCPJSONPermissions(t *testing.T, output, phase string, want []string) {
	t.Helper()
	data := []byte(output)
	if phase != "" {
		var roles map[string]json.RawMessage
		if err := json.Unmarshal(data, &roles); err != nil {
			t.Fatalf("invalid two-role JSON: %v\n%s", err, output)
		}
		data = roles[phase]
	}
	var role struct {
		IncludedPermissions []string `json:"includedPermissions"`
	}
	if err := json.Unmarshal(data, &role); err != nil {
		t.Fatalf("invalid role JSON: %v\n%s", err, output)
	}
	if len(role.IncludedPermissions) == 0 {
		t.Fatalf("missing GCP permissions: %s", output)
	}
	if want != nil && !reflect.DeepEqual(role.IncludedPermissions, want) {
		t.Errorf("permissions = %v, want %v", role.IncludedPermissions, want)
	}
}
