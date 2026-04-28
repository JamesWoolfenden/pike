resource "azurerm_logic_app_integration_account" "pike_gen" {
  name                = "example-ia"
  resource_group_name = "pike"
  location            = "pike"
  sku_name            = "Standard"
  tags = {
    foo = "bar"
  }
}
