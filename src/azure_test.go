package pike

import (
	"reflect"
	"testing"
)

func TestGetAZUREPermissions(t *testing.T) {
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
			got, err := GetAZUREPermissions(tt.args.result)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAZUREPermissions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAZUREPermissions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAZUREResourcePermissions(t *testing.T) {
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
			got, err := GetAZUREResourcePermissions(tt.args.result)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAZUREResourcePermissions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAZUREResourcePermissions() = %v, want %v", got, tt.want)
			}
		})
	}
}
