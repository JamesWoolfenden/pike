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
		{"pass", args{
			ResourceV2{"data",
				"azurerm_resource_group",
				"pike",
				"azurerm",
				[]string{"name", "location", "tags"}},
		}, []string{"Microsoft.Resources/subscriptions/resourcegroups/read"}, false},
		{"empty", args{}, nil, true},
		{"guff", args{ResourceV2{"data",
			"azurerm_toffee_group",
			"pike",
			"azurerm",
			[]string{"name", "location", "tags"}}},
			nil,
			true},
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
