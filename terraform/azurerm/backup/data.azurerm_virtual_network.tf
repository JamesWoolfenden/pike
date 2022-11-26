data "azurerm_virtual_network" "pike" {
  name                = "pike"
  resource_group_name = data.azurerm_resource_group.pike.name
}
