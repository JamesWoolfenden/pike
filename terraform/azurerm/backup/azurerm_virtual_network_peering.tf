resource "azurerm_virtual_network_peering" "pike" {
  name                      = "peer1to2"
  resource_group_name       = "pike"
  virtual_network_name      = data.azurerm_virtual_network.pike.name
  remote_virtual_network_id = data.azurerm_virtual_network.example.id
}

data "azurerm_virtual_network" "pike" {
  name                = "pike"
  resource_group_name = "pike"
}

data "azurerm_virtual_network" "example" {
  name                = "example-network"
  resource_group_name = "pike"
}
