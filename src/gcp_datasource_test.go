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
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetGCPDataPermissions(tt.args.result); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGCPDataPermissions() = %v, want %v", got, tt.want)
			}
		})
	}
}
