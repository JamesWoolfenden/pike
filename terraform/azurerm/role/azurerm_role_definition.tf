resource "azurerm_role_definition" "example" {
  role_definition_id = local.uuid
  name               = "my-custom-role-definition-pike"
  scope              = data.azurerm_subscription.primary.id

  permissions {
    actions = [
      #analytics
      "Microsoft.OperationsManagement/solutions/read",
      "Microsoft.OperationsManagement/solutions/write",
      "Microsoft.OperationsManagement/solutions/delete",

      "Microsoft.Resources/subscriptions/providers/read",

      #role
      "Microsoft.Authorization/roleAssignments/read",
      "Microsoft.Authorization/roleAssignments/write",
      "Microsoft.Authorization/roleAssignments/delete"
    ]
    not_actions = []
  }

  assignable_scopes = [
    data.azurerm_subscription.primary.id,
  ]
}
