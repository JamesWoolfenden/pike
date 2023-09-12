resource "azurerm_app_service_plan" "pike" {
  resource_group_name = "pike"
  location            = "uk south"
  name                = "pikes"
  sku {
    tier = "Standard"
    size = "S1"
  }
  tags = {
    pike = "permissions"
  }
}
