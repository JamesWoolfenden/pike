package parse

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"sort"
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
)

// Test_locateTerraform verifies that locateTerraform finds a tofu or terraform
// binary when one is present on PATH. This pins the fix that added tofu
// detection: previously only "terraform" was checked, so CI (which installs
// "tofu") fell through to hc-install.
func Test_locateTerraform(t *testing.T) {
	t.Parallel()

	hasBinary := false
	for _, bin := range []string{"tofu", "terraform"} {
		if p, err := exec.LookPath(bin); err == nil && p != "" {
			hasBinary = true
			break
		}
	}

	if !hasBinary {
		t.Skip("neither tofu nor terraform on PATH")
	}

	got, err := locateTerraform()
	if err != nil {
		t.Fatalf("locateTerraform() error = %v", err)
	}

	if got == "" {
		t.Error("locateTerraform() returned empty path")
	}
}

// TestParse_createsJSON is an integration test that runs the full Parse
// pipeline against a real provider and asserts that the members JSON file
// is created with non-empty resource and datasource lists.
//
// Requires tofu or terraform on PATH plus registry network access; skipped
// when neither binary is present or when -short is set.
func TestParse_createsJSON(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	hasBinary := false
	for _, bin := range []string{"tofu", "terraform"} {
		if p, err := exec.LookPath(bin); err == nil && p != "" {
			hasBinary = true
			break
		}
	}

	if !hasBinary {
		t.Skip("neither tofu nor terraform on PATH")
	}

	providers := []string{"aws", "azurerm", "google"}

	for _, name := range providers {
		t.Run(name, func(t *testing.T) {
			dir := t.TempDir()
			orig, err := os.Getwd()
			if err != nil {
				t.Fatal(err)
			}

			if err := os.Chdir(dir); err != nil {
				t.Fatal(err)
			}

			t.Cleanup(func() { _ = os.Chdir(orig) })

			if err := Parse("", name); err != nil {
				t.Fatalf("Parse(%q) error = %v", name, err)
			}

			out := filepath.Join(dir, name+"-members.json")

			data, err := os.ReadFile(out)
			if err != nil {
				t.Fatalf("%s-members.json not created: %v", name, err)
			}

			var p provider
			if err := json.Unmarshal(data, &p); err != nil {
				t.Fatalf("%s-members.json is not valid JSON: %v", name, err)
			}

			if len(p.Resources) == 0 {
				t.Errorf("%s-members.json has no resources", name)
			}

			if len(p.DataSources) == 0 {
				t.Errorf("%s-members.json has no data sources", name)
			}

			t.Logf("%s: %d resources, %d data sources", name, len(p.Resources), len(p.DataSources))
		})
	}
}

// Test_extractMembers_Deprecated builds a synthetic ProviderSchemas
// object that mixes deprecated and live resources/datasources, then
// checks that the extracted provider struct has the deprecation maps
// populated with the right names and messages.
//
// We build the schemas by hand rather than by shelling out to a real
// terraform binary so the test stays hermetic and fast — the network
// path is exercised by TestParse with -short gating.
func Test_extractMembers_Deprecated(t *testing.T) {
	t.Parallel()

	schemas := &tfjson.ProviderSchemas{
		Schemas: map[string]*tfjson.ProviderSchema{
			// Hostname'd form; matchesSource must strip the prefix.
			"registry.terraform.io/hashicorp/aws": {
				ResourceSchemas: map[string]*tfjson.Schema{
					"aws_live_thing": {
						Block: &tfjson.SchemaBlock{Description: "alive and well"},
					},
					"aws_dead_thing": {
						Block: &tfjson.SchemaBlock{
							Deprecated:  true,
							Description: "Deprecated: use aws_live_thing instead",
						},
					},
					"aws_deprecated_without_message": {
						Block: &tfjson.SchemaBlock{Deprecated: true},
					},
				},
				DataSourceSchemas: map[string]*tfjson.Schema{
					"aws_live_data": {
						Block: &tfjson.SchemaBlock{Description: "still a datasource"},
					},
					"aws_dead_data": {
						Block: &tfjson.SchemaBlock{
							Deprecated:  true,
							Description: "Deprecated: query via aws_live_data",
						},
					},
				},
			},
			// A second provider under the same schemas map must be ignored
			// when we're asking about aws — matches-by-suffix not by
			// membership.
			"registry.terraform.io/hashicorp/random": {
				ResourceSchemas: map[string]*tfjson.Schema{
					"random_id": {Block: &tfjson.SchemaBlock{}},
				},
			},
		},
	}

	got := extractMembers(schemas, "hashicorp/aws")

	wantResources := []string{"aws_dead_thing", "aws_deprecated_without_message", "aws_live_thing"}
	sort.Strings(wantResources)
	if !reflect.DeepEqual(got.Resources, wantResources) {
		t.Errorf("Resources = %v, want %v", got.Resources, wantResources)
	}

	wantData := []string{"aws_dead_data", "aws_live_data"}
	sort.Strings(wantData)
	if !reflect.DeepEqual(got.DataSources, wantData) {
		t.Errorf("DataSources = %v, want %v", got.DataSources, wantData)
	}

	wantDeprecatedResources := map[string]string{
		"aws_dead_thing":                 "Deprecated: use aws_live_thing instead",
		"aws_deprecated_without_message": "",
	}
	if !reflect.DeepEqual(got.DeprecatedResources, wantDeprecatedResources) {
		t.Errorf("DeprecatedResources = %v, want %v", got.DeprecatedResources, wantDeprecatedResources)
	}

	wantDeprecatedData := map[string]string{
		"aws_dead_data": "Deprecated: query via aws_live_data",
	}
	if !reflect.DeepEqual(got.DeprecatedData, wantDeprecatedData) {
		t.Errorf("DeprecatedData = %v, want %v", got.DeprecatedData, wantDeprecatedData)
	}
}

// Test_extractMembers_NoDeprecations covers the expected-common case: a
// provider with no deprecated members must leave the two maps nil so the
// serialised JSON omits them entirely (`omitempty`).
func Test_extractMembers_NoDeprecations(t *testing.T) {
	t.Parallel()

	schemas := &tfjson.ProviderSchemas{
		Schemas: map[string]*tfjson.ProviderSchema{
			"hashicorp/google": {
				ResourceSchemas: map[string]*tfjson.Schema{
					"google_storage_bucket": {Block: &tfjson.SchemaBlock{}},
				},
				DataSourceSchemas: map[string]*tfjson.Schema{
					"google_storage_bucket": {Block: &tfjson.SchemaBlock{}},
				},
			},
		},
	}

	got := extractMembers(schemas, "hashicorp/google")

	if got.DeprecatedResources != nil {
		t.Errorf("DeprecatedResources = %v, want nil", got.DeprecatedResources)
	}
	if got.DeprecatedData != nil {
		t.Errorf("DeprecatedData = %v, want nil", got.DeprecatedData)
	}
}
