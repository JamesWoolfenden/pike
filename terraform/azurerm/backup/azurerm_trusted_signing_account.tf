resource "azurerm_trusted_signing_account" "pike_gen" {
  name                = "example-account"
  resource_group_name = "pike"
  location            = "West Europe"
  sku_name            = "Basic"
}
