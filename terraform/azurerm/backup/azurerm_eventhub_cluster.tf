resource "azurerm_eventhub_cluster" "pike_gen" {
  name                = "example"
  resource_group_name = "pike"
  location            = "pike"
  sku_name            = "Dedicated_1"
}
