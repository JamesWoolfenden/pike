package pike

import (
	"reflect"
	"testing"
)

func TestGetAZUREDataPermissions(t *testing.T) {
	type args struct {
		result ResourceV2
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAZUREDataPermissions(tt.args.result)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAZUREDataPermissions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAZUREDataPermissions() = %v, want %v", got, tt.want)
			}
		})
	}
}
