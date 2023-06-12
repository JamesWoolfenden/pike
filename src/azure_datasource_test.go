package pike_test

import (
	"reflect"
	"testing"

	pike "github.com/jameswoolfenden/pike/src"
)

func TestGetAZUREDataPermissions(t *testing.T) {
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
		{"pass", args{
			pike.ResourceV2{
				TypeName:     "data",
				Name:         "azurerm_resource_group",
				ResourceName: "pike",
				Provider:     "azurerm",
				Attributes:   []string{"name", "location", "tags"},
			},
		}, []string{"Microsoft.Resources/subscriptions/resourcegroups/read"}, false},
		{"empty", args{}, nil, true},
		{
			"guff",
			args{pike.ResourceV2{
				TypeName:     "data",
				Name:         "azurerm_toffee_group",
				ResourceName: "pike",
				Provider:     "azurerm",
				Attributes:   []string{"name", "location", "tags"},
			}},
			nil,
			true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := pike.GetAZUREDataPermissions(tt.args.result)
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
