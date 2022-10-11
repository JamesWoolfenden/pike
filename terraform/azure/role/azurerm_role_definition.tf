resource "azurerm_role_definition" "example" {
  role_definition_id = local.uuid
  name               = "my-custom-role-definition-pike"
  scope              = data.azurerm_subscription.primary.id

  permissions {
    actions = [
      "Microsoft.Web/serverfarms/read",
      "Microsoft.Web/serverfarms/write",
      "Microsoft.Web/serverfarms/delete",
      "Microsoft.Resources/subscriptions/providers/read"
    ]
    not_actions = []
  }

  assignable_scopes = [
    data.azurerm_subscription.primary.id,
  ]
}
