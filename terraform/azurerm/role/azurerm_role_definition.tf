resource "azurerm_role_definition" "example" {
  role_definition_id = local.uuid
  name               = "my-custom-role-definition-pike"
  scope              = data.azurerm_subscription.primary.id

  permissions {
    actions = [
      "Microsoft.OperationalInsights/workspaces/read",
      "Microsoft.OperationalInsights/workspaces/write",
      "Microsoft.OperationalInsights/workspaces/delete",

      "Microsoft.Storage/storageAccounts/read",
      "Microsoft.Storage/storageAccounts/write",
      "Microsoft.Storage/storageAccounts/delete",
      "Microsoft.Storage/storageAccounts/listKeys/action",
      "Microsoft.Storage/storageAccounts/blobServices/read",
      "Microsoft.Storage/storageAccounts/fileServices/read",

      "Microsoft.Network/networkWatchers/flowLogs/read",
      "Microsoft.Network/networkWatchers/flowLogs/write",
      "Microsoft.Network/networkWatchers/flowLogs/delete",
      "Microsoft.Network/networkWatchers/read",
      "Microsoft.Network/networkSecurityGroups/write",
      "Microsoft.OperationalInsights/workspaces/sharedKeys/action",


      "Microsoft.Network/applicationGateways/read",
      "Microsoft.Network/connections/read",
      "Microsoft.Network/loadBalancers/read",
      "Microsoft.Network/localNetworkGateways/read",
      "Microsoft.Network/networkInterfaces/read",
      "Microsoft.Network/networkSecurityGroups/read",
      "Microsoft.Network/publicIPAddresses/read",
      "Microsoft.Network/routeTables/read",
      "Microsoft.Network/virtualNetworkGateways/read",
      "Microsoft.Network/virtualNetworks/read",
      "Microsoft.Network/expressRouteCircuits/read"

    ]
    not_actions = []
  }

  assignable_scopes = [
    data.azurerm_subscription.primary.id,
  ]
}
