resource "azurerm_role_definition" "example" {
  role_definition_id = local.uuid
  name               = "my-custom-role-definition-pike"
  scope              = data.azurerm_subscription.primary.id

  permissions {
    actions = [
      "Microsoft.DBforMariaDB/servers/read",
      "Microsoft.DBforMariaDB/servers/write",
      "Microsoft.DBforMariaDB/servers/delete",
    ]
    not_actions = []
  }

  assignable_scopes = [
    data.azurerm_subscription.primary.id,
  ]
}
