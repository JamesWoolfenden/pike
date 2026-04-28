resource "azurerm_network_function_azure_traffic_collector" "pike_gen" {
  name                = "example-nfatc"
  resource_group_name = "pike"
  location            = "West US"

  tags = {
    key = "value"
  }
}
