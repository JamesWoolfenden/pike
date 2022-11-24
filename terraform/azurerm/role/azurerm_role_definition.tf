resource "azurerm_role_definition" "example" {
  role_definition_id = local.uuid
  name               = "my-custom-role-definition-pike"
  scope              = data.azurerm_subscription.primary.id

  permissions {
    actions = [

      //data azurerm_key_vault
      "Microsoft.Resources/subscriptions/resourcegroups/read",
      "Microsoft.KeyVault/vaults/read",

      // data.azurerm_key_vault_key
      "Microsoft.KeyVault/vaults/keys/read",
      "Microsoft.KeyVault/vaults/keys/versions/read",

      //user_identity
      "Microsoft.Resources/subscriptions/resourcegroups/read",
      "Microsoft.ManagedIdentity/userAssignedIdentities/read",
      "Microsoft.ManagedIdentity/userAssignedIdentities/write",
      "Microsoft.ManagedIdentity/userAssignedIdentities/delete",
      //ACR
      "Microsoft.Resources/subscriptions/resourcegroups/read",
      "Microsoft.ManagedIdentity/userAssignedIdentities/read",
      "Microsoft.ContainerRegistry/registries/read",
      "Microsoft.ContainerRegistry/registries/write",
      "Microsoft.ManagedIdentity/userAssignedIdentities/assign/action",
      "Microsoft.ContainerRegistry/registries/operationStatuses/read",
      "Microsoft.ContainerRegistry/registries/replications/write",
      "Microsoft.ContainerRegistry/registries/replications/operationStatuses/read",
      "Microsoft.ContainerRegistry/registries/delete"
    ]
    not_actions = []
  }

  assignable_scopes = [
    data.azurerm_subscription.primary.id,
  ]
}
