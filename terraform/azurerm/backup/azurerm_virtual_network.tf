resource "azurerm_virtual_network" "pike" {
  name                = "example-network"
  resource_group_name = "pike"
  location            = "uksouth"
  address_space       = ["10.0.0.0/16"]
}
