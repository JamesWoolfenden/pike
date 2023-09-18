resource "azurerm_role_definition" "example" {
  role_definition_id = local.uuid
  name               = "my-custom-role-definition-pike"
  scope              = data.azurerm_subscription.primary.id

  permissions {
    actions = [
      "Microsoft.Resources/subscriptions/providers/read",
      //data.azurerm_storage_management_policy
      //      "Microsoft.Storage/storageAccounts/read",
      //azurerm_storage_encryption_scope
      #      "Microsoft.Storage/storageAccounts/encryptionScopes/read",
      #      //azurerm_storage_table_entity, azurerm_storage_share, azurerm_storage_container, azurerm_storage_blob
      #      "Microsoft.Storage/storageAccounts/listKeys/action",

    ]
    not_actions = []
  }

  assignable_scopes = [
    data.azurerm_subscription.primary.id,
  ]
}
