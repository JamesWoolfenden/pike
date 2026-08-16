package pike

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestExtractServiceAccount(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		attributes []string
		want       string
	}{
		{"with service_account", []string{"name", "service_account", "location"}, "custom"},
		{"without service_account", []string{"name", "location"}, "default"},
		{"empty attributes", []string{}, "default"},
		{"nil attributes", nil, "default"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := extractServiceAccount(tt.attributes); got != tt.want {
				t.Errorf("extractServiceAccount() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestGCPLookup(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		resource string
		wantNil  bool
	}{
		{"known resource", "google_compute_instance", false},
		{"unknown resource", "google_does_not_exist_xyz", true},
		{"empty name", "", true},
		{"iam member variant strips suffix", "google_project_iam_member", false},
		{"iam binding variant strips suffix", "google_project_iam_binding", false},
		{"iam policy variant strips suffix", "google_project_iam_policy", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := GCPLookup(tt.resource)
			if tt.wantNil && got != nil {
				t.Errorf("GCPLookup(%q) = non-nil, want nil", tt.resource)
			}
			if !tt.wantNil && got == nil {
				t.Errorf("GCPLookup(%q) = nil, want non-nil", tt.resource)
			}
		})
	}
}

func TestGetGCPPermissions(t *testing.T) {
	t.Parallel()

	type args struct {
		result ResourceV2
	}

	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{name: "missing", args: args{
			result: ResourceV2{
				TypeName: "bogus", Name: "bogus",
			},
		}, wantErr: true},
		{name: "notype", args: args{
			result: ResourceV2{
				TypeName: "bogus", Name: "google_compute_duff", ResourceName: "pike", Provider: "azurerm",
				Attributes: []string{
					"name",
					"machine_type", "zone",
				},
			},
		}, wantErr: true},
		{name: "not implemented", args: args{
			result: ResourceV2{
				TypeName: "data", Name: "google_compute_duff", ResourceName: "pike", Provider: "azurerm",
				Attributes: []string{
					"name",
					"machine_type", "zone",
				},
			},
		}, wantErr: true},
		{
			name: "resource",
			args: args{
				result: ResourceV2{
					TypeName: "resource", Name: "google_compute_instance",
					Attributes: []string{"name", "machine_type", "zone"},
				},
			},
			want: []string{
				"compute.zones.get",
				"compute.instances.create",
				"compute.instances.get",
				"compute.disks.create",
				"compute.disks.create",
				"compute.subnetworks.use",
				"compute.subnetworks.useExternalIp",
				"compute.instances.setMetadata",
				"compute.instances.delete",
				"compute.instances.setLabels",
				"compute.disks.setLabels",
				"compute.instances.updateNetworkInterface",
				"compute.instances.get",
				"compute.zones.get",
				"compute.instances.delete",
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := getGCPPermissions(tt.args.result)

			if (err != nil) != tt.wantErr {
				t.Errorf("getGCPPermissions() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getGCPPermissions() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestGetGCPAppEnginePermissions covers the App Engine resource mappings
// (mapping/google/resource/appengine/*.json) that previously shipped as
// empty stubs -- apply/destroy/modify/plan all [] -- so any Terraform using
// these resource types silently floored to zero required IAM permissions.
// Verified against the official App Engine Admin API "Authorization
// requires..." IAM permission for each REST method (apps.services.versions,
// apps.domainMappings, apps.firewall.ingressRules).
func TestGetGCPAppEnginePermissions(t *testing.T) {
	t.Parallel()

	type args struct {
		result ResourceV2
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "standard app version",
			args: args{
				result: ResourceV2{
					TypeName: "resource", Name: "google_app_engine_standard_app_version",
					Attributes: []string{"service", "runtime"},
				},
			},
			want: []string{
				"appengine.versions.create",
				"appengine.versions.get",
				"appengine.versions.update",
				"appengine.versions.delete",
				"appengine.versions.get",
				"appengine.versions.update",
				"appengine.versions.delete",
			},
		},
		{
			name: "flexible app version",
			args: args{
				result: ResourceV2{
					TypeName: "resource", Name: "google_app_engine_flexible_app_version",
					Attributes: []string{"service", "runtime"},
				},
			},
			want: []string{
				"appengine.versions.create",
				"appengine.versions.get",
				"appengine.versions.update",
				"appengine.versions.delete",
				"appengine.versions.get",
				"appengine.versions.update",
				"appengine.versions.delete",
			},
		},
		{
			name: "domain mapping",
			args: args{
				result: ResourceV2{
					TypeName: "resource", Name: "google_app_engine_domain_mapping",
					Attributes: []string{"domain_name"},
				},
			},
			want: []string{
				"appengine.applications.get",
				"appengine.applications.update",
				"appengine.applications.get",
				"appengine.applications.update",
				"appengine.applications.update",
			},
		},
		{
			name: "firewall rule",
			args: args{
				result: ResourceV2{
					TypeName: "resource", Name: "google_app_engine_firewall_rule",
					Attributes: []string{"priority", "action"},
				},
			},
			want: []string{
				"appengine.applications.get",
				"appengine.applications.update",
				"appengine.applications.get",
				"appengine.applications.update",
				"appengine.applications.update",
			},
		},
		{
			name: "application url dispatch rules",
			args: args{
				result: ResourceV2{
					TypeName: "resource", Name: "google_app_engine_application_url_dispatch_rules",
					Attributes: []string{"dispatch_rules"},
				},
			},
			want: []string{
				"appengine.applications.get",
				"appengine.applications.update",
				"appengine.applications.get",
				"appengine.applications.update",
			},
		},
		{
			name: "service split traffic",
			args: args{
				result: ResourceV2{
					TypeName: "resource", Name: "google_app_engine_service_split_traffic",
					Attributes: []string{"service", "split"},
				},
			},
			want: []string{
				"appengine.services.update",
				"appengine.services.get",
				"appengine.services.get",
				"appengine.services.update",
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := getGCPPermissions(tt.args.result)
			if err != nil {
				t.Fatalf("getGCPPermissions() unexpected error = %v", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getGCPPermissions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetGCPResourcePermissions(t *testing.T) {
	t.Parallel()

	type args struct {
		result ResourceV2
	}

	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{name: "missing", args: args{result: ResourceV2{TypeName: "bogus", Name: "bogus"}}},
		{
			name: "resource",
			args: args{result: ResourceV2{TypeName: "resource", Name: "google_compute_instance", Attributes: []string{
				"name",
				"machine_type", "zone",
			}}},
			want: []string{
				"compute.zones.get",
				"compute.instances.create",
				"compute.instances.get",
				"compute.disks.create",
				"compute.disks.create",
				"compute.subnetworks.use",
				"compute.subnetworks.useExternalIp",
				"compute.instances.setMetadata",
				"compute.instances.delete",
				"compute.instances.setLabels",
				"compute.disks.setLabels",
				"compute.instances.updateNetworkInterface",
				"compute.instances.get",
				"compute.zones.get",
				"compute.instances.delete",
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got, _ := getGCPResourcePermissions(tt.args.result); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getGCPResourcePermissions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInvalidGCPResourceError(t *testing.T) {
	tests := []struct {
		name     string
		resource string
		want     string
	}{
		{
			name:     "empty resource",
			resource: "",
			want:     "Invalid GCP lookup sourceData type for resource ",
		},
		{
			name:     "valid resource",
			resource: "google_storage_bucket",
			want:     "Invalid GCP lookup sourceData type for resource google_storage_bucket",
		},
		{
			name:     "special characters",
			resource: "test*&^%",
			want:     "Invalid GCP lookup sourceData type for resource test*&^%",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := invalidGCPResourceError{resource: tt.resource}
			if got := err.Error(); got != tt.want {
				t.Errorf("invalidGCPResourceError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInvalidPermissionMapError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		err  error
		want string
	}{
		{
			name: "nil error",
			err:  nil,
			want: "Invalid Permission Map <nil>",
		},
		{
			name: "simple error",
			err:  errors.New("permission denied"),
			want: "Invalid Permission Map permission denied",
		},
		{
			name: "wrapped error",
			err:  fmt.Errorf("wrapped: %w", errors.New("inner error")),
			want: "Invalid Permission Map wrapped: inner error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := invalidPermissionMapError{err: tt.err}

			if got := err.Error(); got != tt.want {
				t.Errorf("invalidPermissionMapError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
