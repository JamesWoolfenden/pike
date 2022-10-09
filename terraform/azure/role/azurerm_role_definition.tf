resource "azurerm_role_definition" "example" {
  role_definition_id = local.uuid
  name               = "my-custom-role-definition-pike"
  scope              = data.azurerm_subscription.primary.id

  permissions {
    actions = [
      "Microsoft.Resources/subscriptions/providers/read",

      "Microsoft.Resources/subscriptions/resourcegroups/read",
      "Microsoft.Resources/subscriptions/resourcegroups/write",

    "Microsoft.Resources/subscriptions/resourcegroups/delete"]
    not_actions = []
  }

  assignable_scopes = [
    data.azurerm_subscription.primary.id,
  ]
}
