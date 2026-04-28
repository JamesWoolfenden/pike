resource "azurerm_servicebus_namespace" "pike_gen" {
  name                = "tfex-servicebus-namespace"
  location            = "pike"
  resource_group_name = "pike"
  sku                 = "Standard"

  tags = {
    source = "terraform"
  }
}
