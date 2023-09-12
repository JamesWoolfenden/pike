resource "azurerm_app_service" "pike" {
  resource_group_name = "pike"
  location            = "uk south"
  name                = "pike"
  app_service_plan_id = azurerm_app_service_plan.pike.id
}
