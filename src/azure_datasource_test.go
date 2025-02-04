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
		{
			name: "pass",
			args: args{
				result: pike.ResourceV2{
					TypeName:     "data",
					Name:         "azurerm_resource_group",
					ResourceName: "pike",
					Provider:     "azurerm",
					Attributes:   []string{"name", "location", "tags"},
				},
			},
			want:    []string{"Microsoft.Resources/subscriptions/resourcegroups/read"},
			wantErr: false,
		},
		{
			name:    "empty",
			args:    args{},
			want:    nil,
			wantErr: true,
		},
		{
			name: "guff",
			args: args{result: pike.ResourceV2{
				TypeName:     "data",
				Name:         "azurerm_toffee_group",
				ResourceName: "pike",
				Provider:     "azurerm",
				Attributes:   []string{"name", "location", "tags"},
			}},
			wantErr: true,
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
