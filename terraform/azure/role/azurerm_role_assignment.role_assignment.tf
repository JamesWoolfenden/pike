locals {
  uuid = uuid()
}
resource "azurerm_role_assignment" "pike" {
  name               = local.uuid
  scope              = data.azurerm_subscription.primary.id
  role_definition_id = azurerm_role_definition.example.role_definition_resource_id
  principal_id       = azuread_service_principal.pike.id
}

data "azurerm_subscription" "primary" {
}
