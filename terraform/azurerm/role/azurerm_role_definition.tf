resource "azurerm_role_definition" "example" {
  role_definition_id = local.uuid
  name               = "my-custom-role-definition-pike"
  scope              = data.azurerm_subscription.primary.id

  permissions {
    actions = [
      "Microsoft.Resources/subscriptions/resourcegroups/read",

      //data azurerm_key_vault
      "Microsoft.Resources/subscriptions/resourcegroups/read",
      "Microsoft.KeyVault/vaults/read",

      // data.azurerm_key_vault_key
      "Microsoft.KeyVault/vaults/keys/read",
      "Microsoft.KeyVault/vaults/keys/versions/read"
    ]
    not_actions = []
  }

  assignable_scopes = [
    data.azurerm_subscription.primary.id,
  ]
}
