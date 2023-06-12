package pike_test

import (
	"reflect"
	"testing"

	pike "github.com/jameswoolfenden/pike/src"
)

func TestGetGCPPermissions(t *testing.T) {
	type args struct {
		result pike.ResourceV2
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"missing", args{pike.ResourceV2{"bogus", "bogus", "", "", nil}}, nil, true},
		{"notype", args{pike.ResourceV2{"bogus", "google_compute_instance", "pike", "azurerm", []string{
			"name",
			"machine_type", "zone",
		}}}, nil, true},
		{"not implemented", args{pike.ResourceV2{"data", "google_compute_instance", "pike", "azurerm", []string{
			"name",
			"machine_type", "zone",
		}}}, nil, true},
		{
			"resource",
			args{
				pike.ResourceV2{
					"resource", "google_compute_instance", "", "",
					[]string{"name", "machine_type", "zone"},
				},
			},
			[]string{
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
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
	type args struct {
		result pike.ResourceV2
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"missing", args{pike.ResourceV2{"bogus", "bogus", "", "", nil}}, nil, false},
		{
			"resource",
			args{pike.ResourceV2{"resource", "google_compute_instance", "", "", []string{
				"name",
				"machine_type", "zone",
			}}},
			[]string{
				"compute.zones.get",
				"compute.instances.create",
				"compute.instances.get",
				"compute.disks.create",
				"compute.disks.create",
				"compute.subnetworks.use",
				"compute.subnetworks.useExternalIp",
				"compute.instances.setMetadata",
				"compute.instances.delete",
				"compute.instances.delete"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := pike.GetGCPResourcePermissions(tt.args.result); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGCPResourcePermissions() = %v, want %v", got, tt.want)
			}
		})
	}
}
