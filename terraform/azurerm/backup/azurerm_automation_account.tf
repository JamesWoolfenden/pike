resource "azurerm_automation_account" "pike_gen" {
  name                = "example-account"
  location            = "pike"
  resource_group_name = "pike"
  sku_name            = "Basic"

  tags = {
    environment = "development"
  }
}
