resource "azurerm_role_definition" "example" {
  role_definition_id = local.uuid
  name               = "my-custom-role-definition-pike"
  scope              = data.azurerm_subscription.primary.id

  permissions {
    actions = [

      //vnet
      #      "Microsoft.Resources/subscriptions/resourcegroups/read",
      #      "Microsoft.Network/virtualNetworks/read",
      #      "Microsoft.Network/virtualNetworks/write",
      #      "Microsoft.Network/virtualNetworks/delete",

      //subnet
      "Microsoft.Resources/subscriptions/resourcegroups/read",
      "Microsoft.Network/virtualNetworks/read",
      "Microsoft.Network/virtualNetworks/subnets/read",
      "Microsoft.Network/virtualNetworks/subnets/write",
      "Microsoft.Network/virtualNetworks/subnets/delete",

      "Microsoft.Resources/subscriptions/resourcegroups/read",
      "Microsoft.Network/virtualNetworks/subnets/read",
      "Microsoft.Compute/virtualMachineScaleSets/read",
      "Microsoft.Compute/virtualMachineScaleSets/write",
      "Microsoft.Network/virtualNetworks/subnets/join/action",
      "Microsoft.Compute/virtualMachineScaleSets/delete"

    ]
    not_actions = []
  }

  assignable_scopes = [
    data.azurerm_subscription.primary.id,
  ]
}
