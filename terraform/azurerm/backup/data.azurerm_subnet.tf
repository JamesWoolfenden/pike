data "azurerm_subnet" "pike" {
  name                 = "pike"
  resource_group_name  = data.azurerm_resource_group.pike.name
  virtual_network_name = data.azurerm_virtual_network.pike.name
}
