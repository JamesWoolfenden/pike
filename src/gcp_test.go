package pike

import (
	"reflect"
	"testing"
)

func TestGetGCPPermissions(t *testing.T) {
	type args struct {
		result ResourceV2
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"missing", args{ResourceV2{"bogus", "bogus", "", "", nil}}, nil},
		{"notype", args{ResourceV2{"bogus", "google_compute_instance", "", "", []string{"name",
			"machine_type", "zone"}}}, nil},
		{"data", args{ResourceV2{"data", "google_compute_instance", "", "", []string{"name",
			"machine_type", "zone"}}}, nil},
		{"resource", args{ResourceV2{"resource", "google_compute_instance", "", "", []string{"name",
			"machine_type", "zone"}}}, []string{"compute.zones.get", "compute.instances.create", "compute.instances.get", "compute.disks.create", "compute.disks.create", "compute.subnetworks.use", "compute.subnetworks.useExternalIp", "compute.instances.setMetadata", "compute.instances.delete", "compute.instances.delete"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetGCPPermissions(tt.args.result); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGCPPermissions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetGCPResourcePermissions(t *testing.T) {
	type args struct {
		result ResourceV2
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"missing", args{ResourceV2{"bogus", "bogus", "", "", nil}}, nil},
		{"resource", args{ResourceV2{"resource", "google_compute_instance", "", "", []string{"name",
			"machine_type", "zone"}}}, []string{"compute.zones.get", "compute.instances.create", "compute.instances.get", "compute.disks.create", "compute.disks.create", "compute.subnetworks.use", "compute.subnetworks.useExternalIp", "compute.instances.setMetadata", "compute.instances.delete", "compute.instances.delete"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetGCPResourcePermissions(tt.args.result); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGCPResourcePermissions() = %v, want %v", got, tt.want)
			}
		})
	}
}
