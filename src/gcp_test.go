package pike_test

import (
	"reflect"
	"testing"

	pike "github.com/jameswoolfenden/pike/src"
)

func TestGetGCPPermissions(t *testing.T) {
	t.Parallel()

	type args struct {
		result pike.ResourceV2
	}

	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{name: "missing", args: args{
			result: pike.ResourceV2{
				TypeName: "bogus", Name: "bogus",
			},
		}, wantErr: true},
		{name: "notype", args: args{
			result: pike.ResourceV2{
				TypeName: "bogus", Name: "google_compute_duff", ResourceName: "pike", Provider: "azurerm", Attributes: []string{
					"name",
					"machine_type", "zone",
				},
			},
		}, wantErr: true},
		{name: "not implemented", args: args{
			result: pike.ResourceV2{
				TypeName: "data", Name: "google_compute_duff", ResourceName: "pike", Provider: "azurerm", Attributes: []string{
					"name",
					"machine_type", "zone",
				},
			},
		}, wantErr: true},
		{
			name: "resource",
			args: args{
				result: pike.ResourceV2{
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
				"compute.instances.delete",
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := pike.GetGCPPermissions(tt.args.result)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGCPPermissions() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGCPPermissions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetGCPResourcePermissions(t *testing.T) {
	t.Parallel()

	type args struct {
		result pike.ResourceV2
	}

	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{name: "missing", args: args{result: pike.ResourceV2{TypeName: "bogus", Name: "bogus"}}},
		{
			name: "resource",
			args: args{result: pike.ResourceV2{TypeName: "resource", Name: "google_compute_instance", Attributes: []string{
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
				"compute.instances.delete",
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got, _ := pike.GetGCPResourcePermissions(tt.args.result); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGCPResourcePermissions() = %v, want %v", got, tt.want)
			}
		})
	}
}
