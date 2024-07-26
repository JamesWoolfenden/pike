package pike_test

import (
	"reflect"
	"testing"

	pike "github.com/jameswoolfenden/pike/src"
)

func TestGetGCPDataPermissions(t *testing.T) {
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
		{"placeholder", args{pike.ResourceV2{
			TypeName:     "data",
			Name:         "google_compute_image",
			ResourceName: "image",
			Provider:     "google",
			Attributes:   []string{"family", "project"},
		}}, nil, false},
		{"basic", args{pike.ResourceV2{
			TypeName:     "data",
			Name:         "google_compute_network",
			ResourceName: "image",
			Provider:     "google",
			Attributes:   []string{"name"},
		}}, []string{"compute.networks.get"}, false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got, _ := pike.GetGCPDataPermissions(tt.args.result); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGCPDataPermissions() = %v, want %v", got, tt.want)
			}
		})
	}
}
