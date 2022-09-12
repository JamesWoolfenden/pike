package pike

import (
	"reflect"
	"testing"
)

func TestGetGCPDataPermissions(t *testing.T) {
	type args struct {
		result ResourceV2
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"placeholder", args{ResourceV2{
			TypeName:     "data",
			Name:         "google_compute_image",
			ResourceName: "image",
			Provider:     "google",
			Attributes:   []string{"family", "project"},
		}}, nil, false},
		{"basic", args{ResourceV2{
			TypeName:     "data",
			Name:         "google_compute_network",
			ResourceName: "image",
			Provider:     "google",
			Attributes:   []string{"name"},
		}}, []string{"compute.networks.get"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := GetGCPDataPermissions(tt.args.result); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGCPDataPermissions() = %v, want %v", got, tt.want)
			}
		})
	}
}
