resource "azurerm_healthbot" "pike_gen" {
  name                = "example-bot"
  resource_group_name = "pike"
  location            = "pike"
  sku_name            = "F0"
}
