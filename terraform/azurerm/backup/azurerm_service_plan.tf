resource "azurerm_service_plan" "pike" {
  location            = "uk south"
  name                = "pike"
  resource_group_name = "pike"
  os_type             = "Linux"
  sku_name            = "P1v2"

  tags = {
    pike = "permissions"
  }
}
