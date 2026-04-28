resource "azurerm_ai_services" "pike_gen" {
  name                = "example-account"
  location            = "pike"
  resource_group_name = "pike"
  sku_name            = "S0"

  tags = {
    Acceptance = "Test"
  }
}
