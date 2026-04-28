data "azurerm_local_network_gateway" "pike_gen" {
  name                = "existing-local-network-gateway"
  resource_group_name = "existing-resources"
}
