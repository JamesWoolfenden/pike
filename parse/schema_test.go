package parse

import (
	"reflect"
	"sort"
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
)

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
