resource "azurerm_relay_namespace" "pike_gen" {
  name                = "example-relay"
  location            = "pike"
  resource_group_name = "pike"

  sku_name = "Standard"

  tags = {
    source = "terraform"
  }
}
