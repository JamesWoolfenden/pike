resource "azurerm_lighthouse_assignment" "pike_gen" {
  scope                    = "pike"
  lighthouse_definition_id = "/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.ManagedServices/registrationDefinitions/00000000-0000-0000-0000-000000000000"
}
