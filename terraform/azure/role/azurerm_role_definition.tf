resource "azurerm_role_definition" "example" {
  role_definition_id = local.uuid
  name               = "my-custom-role-definition-pike"
  scope              = data.azurerm_subscription.primary.id

  permissions {
    actions = [
      "Microsoft.Resources/subscriptions/resourcegroups/read",
      "Microsoft.KeyVault/vaults/read",
      "Microsoft.KeyVault/vaults/write",
      "Microsoft.KeyVault/vaults/delete",
      "Microsoft.KeyVault/locations/deletedVaults/read",

      "Microsoft.DocumentDB/databaseAccounts/read",
      "Microsoft.DocumentDB/databaseAccounts/write",
      "Microsoft.DocumentDB/databaseAccounts/delete",
      "Microsoft.DocumentDB/databaseAccounts/listKeys/action",
      "Microsoft.DocumentDB/databaseAccounts/readonlykeys/action",
      "Microsoft.DocumentDB/databaseAccounts/listConnectionStrings/action",

      "Microsoft.DocumentDB/databaseAccounts/tables/read",
      "Microsoft.DocumentDB/databaseAccounts/tables/write",
      "Microsoft.DocumentDB/databaseAccounts/tables/delete",

    ]
    not_actions = []
  }

  assignable_scopes = [
    data.azurerm_subscription.primary.id,
  ]
}
