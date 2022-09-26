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
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"missing", args{ResourceV2{"bogus", "bogus", "", "", nil}}, nil, false},
		{"notype", args{ResourceV2{"bogus", "google_compute_instance", "", "", []string{"name",
			"machine_type", "zone"}}}, nil, false},
		{"data", args{ResourceV2{"data", "google_compute_instance", "", "", []string{"name",
			"machine_type", "zone"}}}, nil, false},
		{"resource", args{ResourceV2{"resource", "google_compute_instance", "", "", []string{"name",
			"machine_type", "zone"}}}, []string{"compute.zones.get", "compute.instances.create", "compute.instances.get", "compute.disks.create", "compute.disks.create", "compute.subnetworks.use", "compute.subnetworks.useExternalIp", "compute.instances.setMetadata", "compute.instances.delete", "compute.instances.delete"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetGCPPermissions(tt.args.result)
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
		result ResourceV2
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"missing", args{ResourceV2{"bogus", "bogus", "", "", nil}}, nil, false},
		{"resource", args{ResourceV2{"resource", "google_compute_instance", "", "", []string{"name",
			"machine_type", "zone"}}},
			[]string{"compute.zones.get", "compute.instances.create", "compute.instances.get", "compute.disks.create", "compute.disks.create", "compute.subnetworks.use", "compute.subnetworks.useExternalIp", "compute.instances.setMetadata", "compute.instances.delete", "compute.instances.delete"},
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := GetGCPResourcePermissions(tt.args.result); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGCPResourcePermissions() = %v, want %v", got, tt.want)
			}
		})
	}
}
