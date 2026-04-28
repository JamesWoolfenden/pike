resource "azurerm_local_network_gateway" "pike_gen" {
  name                = "backHome"
  resource_group_name = "pike"
  location            = "pike"
  gateway_address     = "12.13.14.15"
  address_space       = ["10.0.0.0/16"]
}
