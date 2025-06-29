package pike

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

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
				"compute.instances.get",
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
				"compute.instances.get",
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
