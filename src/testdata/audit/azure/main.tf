resource "azurerm_role_assignment" "owner" {
  scope                = "/subscriptions/x"
  role_definition_name = "Owner"
  principal_id         = "00000000-0000-0000-0000-000000000000"
}

resource "azurerm_role_assignment" "contributor" {
  scope                = "/subscriptions/x"
  role_definition_name = "Contributor"
  principal_id         = "00000000-0000-0000-0000-000000000000"
}

resource "azurerm_role_definition" "wide" {
  name  = "wide"
  scope = "/subscriptions/x"
  permissions {
    actions     = ["*", "Microsoft.Authorization/roleAssignments/write"]
    not_actions = []
  }
}

resource "azurerm_role_assignment" "ok" {
  scope                = "/subscriptions/x"
  role_definition_name = "Reader"
  principal_id         = "00000000-0000-0000-0000-000000000000"
}
