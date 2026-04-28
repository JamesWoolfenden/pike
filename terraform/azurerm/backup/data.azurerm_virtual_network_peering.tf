data "azurerm_virtual_network_peering" "pike_gen" {
  name               = "peer-vnet01-to-vnet02"
  virtual_network_id = "pike"
}
