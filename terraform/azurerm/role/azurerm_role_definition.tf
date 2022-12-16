resource "azurerm_role_definition" "example" {
  role_definition_id = local.uuid
  name               = "my-custom-role-definition-pike"
  scope              = data.azurerm_subscription.primary.id

  permissions {
    actions = [
      "Microsoft.Security/workspaceSettings/read",
      "Microsoft.Security/workspaceSettings/write",
      "Microsoft.Security/workspaceSettings/delete",

      "Microsoft.Security/securityContacts/read",
      "Microsoft.Security/securityContacts/write",
      "Microsoft.Security/securityContacts/delete",

      "Microsoft.Security/settings/read",
      "Microsoft.Security/settings/write"
    ]
    not_actions = []
  }

  assignable_scopes = [
    data.azurerm_subscription.primary.id,
  ]
}
