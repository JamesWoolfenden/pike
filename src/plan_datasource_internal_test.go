package pike

import (
	"reflect"
	"testing"
)

// These tests guard against the getXPlanPermissions functions looking up
// plan permissions in the resource mapping table for a `data` block instead
// of the data-source mapping table. All three resource names below exist in
// both the resource and data mapping trees with different "plan" contents,
// so a regression (using the resource-side mapping for a data source) shows
// up as a mismatched result rather than merely an error.
func Test_getPlanPermissions_dataSource(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		fn   func(ResourceV2) ([]string, error)
		res  ResourceV2
		want []string
	}{
		{
			name: "aws data source uses data plan permissions, not resource's",
			fn:   getAWSPlanPermissions,
			res: ResourceV2{
				TypeName:   data,
				Name:       "aws_cloudwatch_log_group",
				Attributes: []string{"tags"},
			},
			want: []string{"logs:DescribeLogGroups", "logs:ListTagsLogGroup"},
		},
		{
			name: "gcp data source uses data plan permissions, not resource's",
			fn:   getGCPPlanPermissions,
			res: ResourceV2{
				TypeName:   data,
				Name:       "google_access_context_manager_access_policy",
				Attributes: []string{"tags"},
			},
			want: []string{"accesscontextmanager.accessPolicies.list"},
		},
		{
			name: "azure data source uses data plan permissions, not resource's",
			fn:   getAZUREPlanPermissions,
			res: ResourceV2{
				TypeName:   data,
				Name:       "azurerm_active_directory_domain_service",
				Attributes: []string{"tags"},
			},
			want: []string{"Microsoft.AAD/domainServices/read"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := tt.fn(tt.res)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

// A pure data-only type (no resource-map counterpart) must not error out
// when asked for plan permissions.
func Test_getAWSPlanPermissions_dataOnlyType(t *testing.T) {
	t.Parallel()

	got, err := getAWSPlanPermissions(ResourceV2{
		TypeName:   data,
		Name:       "aws_caller_identity",
		Attributes: nil,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 0 {
		t.Errorf("got %v, want empty", got)
	}
}

// The resource path must keep working unchanged.
func Test_getAWSPlanPermissions_resource(t *testing.T) {
	t.Parallel()

	got, err := getAWSPlanPermissions(ResourceV2{
		TypeName:   resource,
		Name:       "aws_cloudwatch_log_group",
		Attributes: []string{"tags"},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// the resource mapping's "plan" array is empty for this type
	if len(got) != 0 {
		t.Errorf("got %v, want empty (resource plan array is empty)", got)
	}
}
