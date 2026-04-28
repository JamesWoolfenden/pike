resource "azurerm_eventhub_namespace" "pike_gen" {
  name                = "example-namespace"
  location            = "pike"
  resource_group_name = "pike"
  sku                 = "Standard"
  capacity            = 2

  tags = {
    environment = "Production"
  }
}
