data "azurerm_app_service" "pike" {
  resource_group_name = "pike"
  name                = azurerm_app_service.pike.name
}
