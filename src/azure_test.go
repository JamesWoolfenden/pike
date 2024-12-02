package pike_test

import (
	"reflect"
	"testing"

	pike "github.com/jameswoolfenden/pike/src"
)

func TestGetAZUREPermissions(t *testing.T) {
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
			"data",
			args{
				pike.ResourceV2{
					TypeName:     "data",
					Name:         "azurerm_key_vault",
					ResourceName: "MyDemoApiKey",
					Provider:     "azurerm",
					Attributes:   []string{"name"},
				},
			},
			[]string{
				"Microsoft.Resources/subscriptions/resourcegroups/read",
				"Microsoft.KeyVault/vaults/read",
			},
			false,
		},
		{
			"resource",
			args{
				pike.ResourceV2{
					TypeName:     "resource",
					Name:         "azurerm_key_vault",
					ResourceName: "MyDemoApiKey",
					Provider:     "azurerm",
					Attributes:   []string{"name", "name", "resource_group"},
				},
			},
			[]string{
				"Microsoft.Resources/subscriptions/resourcegroups/read",
				"Microsoft.KeyVault/vaults/read",
				"Microsoft.KeyVault/vaults/write",
				"Microsoft.KeyVault/vaults/delete",
				"Microsoft.KeyVault/locations/deletedVaults/read",
			},
			false,
		},
		{
			"bogus resource",
			args{
				pike.ResourceV2{
					TypeName:     "resource",
					Name:         "azurerm_is_best",
					ResourceName: "MyDemoApiKey",
					Provider:     "azurerm",
					Attributes:   []string{"name", "type", "resource_group"},
				},
			},
			nil,
			true,
		},
		{
			"bogus data",
			args{
				pike.ResourceV2{
					TypeName:     "data",
					Name:         "azurerm_is_best",
					ResourceName: "MyDemoApiKey",
					Provider:     "azurerm",
					Attributes:   []string{"name", "type", "resource_group"},
				},
			},
			nil,
			true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := pike.GetAZUREPermissions(tt.args.result)
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
			"resource",
			args{
				pike.ResourceV2{
					TypeName:     "resource",
					Name:         "azurerm_key_vault",
					ResourceName: "MyDemoApiKey",
					Provider:     "azurerm",
					Attributes:   []string{"name", "name", "resource_group"},
				},
			},
			[]string{
				"Microsoft.Resources/subscriptions/resourcegroups/read",
				"Microsoft.KeyVault/vaults/read",
				"Microsoft.KeyVault/vaults/write",
				"Microsoft.KeyVault/vaults/delete",
				"Microsoft.KeyVault/locations/deletedVaults/read",
			},
			false,
		},
		{"bogus", args{pike.ResourceV2{
			TypeName:     "resource",
			Name:         "azurerm_is_best",
			ResourceName: "MyDemoApiKey",
			Provider:     "azurerm",
			Attributes:   []string{"name", "type", "resource_group"},
		}}, nil, true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := pike.GetAZUREResourcePermissions(tt.args.result)

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
