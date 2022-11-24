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
		{"data",
			args{
				ResourceV2{"data",
					"azurerm_key_vault",
					"MyDemoApiKey",
					"azurerm",
					[]string{"name"}}},
			[]string{
				"Microsoft.Resources/subscriptions/resourcegroups/read",
				"Microsoft.KeyVault/vaults/read",
			},
			false},
		{"resource",
			args{
				ResourceV2{"resource",
					"azurerm_key_vault",
					"MyDemoApiKey",
					"azurerm",
					[]string{"name", "name", "resource_group"}}},
			[]string{"Microsoft.Resources/subscriptions/resourcegroups/read",
				"Microsoft.KeyVault/vaults/read",
				"Microsoft.KeyVault/vaults/write",
				"Microsoft.KeyVault/vaults/delete",
				"Microsoft.KeyVault/locations/deletedVaults/read"},
			false},
		{"bogus resource",
			args{
				ResourceV2{"resource",
					"azurerm_is_best",
					"MyDemoApiKey",
					"azurerm",
					[]string{"name", "type", "resource_group"}}},
			nil,
			true},
		{"bogus data",
			args{
				ResourceV2{"data",
					"azurerm_is_best",
					"MyDemoApiKey",
					"azurerm",
					[]string{"name", "type", "resource_group"}}},
			nil,
			true},
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
		{"resource",
			args{
				ResourceV2{"resource",
					"azurerm_key_vault",
					"MyDemoApiKey",
					"azurerm",
					[]string{"name", "name", "resource_group"}}},
			[]string{"Microsoft.Resources/subscriptions/resourcegroups/read",
				"Microsoft.KeyVault/vaults/read",
				"Microsoft.KeyVault/vaults/write",
				"Microsoft.KeyVault/vaults/delete",
				"Microsoft.KeyVault/locations/deletedVaults/read"},
			false},
		{"bogus", args{ResourceV2{"resource",
			"azurerm_is_best",
			"MyDemoApiKey",
			"azurerm",
			[]string{"name", "type", "resource_group"}}}, nil, true},
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
