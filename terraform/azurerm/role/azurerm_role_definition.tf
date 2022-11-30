resource "azurerm_role_definition" "example" {
  role_definition_id = local.uuid
  name               = "my-custom-role-definition-pike"
  scope              = data.azurerm_subscription.primary.id

  permissions {
    actions = [
      "Microsoft.Cache/redis/read",
      "Microsoft.Cache/redis/write",
      "Microsoft.Cache/redis/delete",
      "Microsoft.Cache/redis/listKeys/action",
      "Microsoft.Cache/redis/patchSchedules/delete",

      "Microsoft.Network/privateEndpoints/read",
      "Microsoft.Network/privateEndpoints/write",
      "Microsoft.Network/privateEndpoints/delete",
      "Microsoft.Network/virtualNetworks/subnets/join/action",
      "Microsoft.Cache/redis/PrivateEndpointConnectionsApproval/action",
      //private_dns_zone_group
      "Microsoft.Network/privateEndpoints/privateDnsZoneGroups/write",
      "Microsoft.Network/privateDnsZones/join/action"

    ]
    not_actions = []
  }

  assignable_scopes = [
    data.azurerm_subscription.primary.id,
  ]
}
