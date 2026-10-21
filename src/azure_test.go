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
				"Microsoft.KeyVault/vaults/read",
				"Microsoft.Resources/subscriptions/resourcegroups/read",
				"Microsoft.KeyVault/vaults/write",
				"Microsoft.KeyVault/locations/deletedVaults/read",
				"Microsoft.KeyVault/vaults/delete",
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
				"Microsoft.KeyVault/vaults/read",
				"Microsoft.Resources/subscriptions/resourcegroups/read",
				"Microsoft.KeyVault/vaults/write",
				"Microsoft.KeyVault/locations/deletedVaults/read",
				"Microsoft.KeyVault/vaults/delete",
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
		{
			"private_endpoint redis target gets the redis approval permission",
			args{
				pike.ResourceV2{
					TypeName:        "resource",
					Name:            "azurerm_private_endpoint",
					ResourceName:    "pike",
					Provider:        "azurerm",
					Attributes:      []string{"private_service_connection", "subresource_names"},
					AttributeValues: map[string][]string{"subresource_names": {"redisCache"}},
				},
			},
			[]string{
				"Microsoft.Network/privateEndpoints/read",
				"Microsoft.Network/privateEndpoints/write",
				"Microsoft.Network/privateEndpoints/delete",
				"Microsoft.Network/virtualNetworks/subnets/join/action",
				"Microsoft.Network/privateEndpoints/read",
				"Microsoft.Network/privateEndpoints/write",
				"Microsoft.Network/privateEndpoints/delete",
				"Microsoft.Cache/redis/PrivateEndpointConnectionsApproval/action",
			},
			false,
		},
		{
			"private_endpoint non-redis target gets no guessed approval permission",
			args{
				pike.ResourceV2{
					TypeName:        "resource",
					Name:            "azurerm_private_endpoint",
					ResourceName:    "pike",
					Provider:        "azurerm",
					Attributes:      []string{"private_service_connection", "subresource_names"},
					AttributeValues: map[string][]string{"subresource_names": {"vault"}},
				},
			},
			[]string{
				"Microsoft.Network/privateEndpoints/read",
				"Microsoft.Network/privateEndpoints/write",
				"Microsoft.Network/privateEndpoints/delete",
				"Microsoft.Network/virtualNetworks/subnets/join/action",
				"Microsoft.Network/privateEndpoints/read",
				"Microsoft.Network/privateEndpoints/write",
				"Microsoft.Network/privateEndpoints/delete",
			},
			false,
		},
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
