data "azurerm_virtual_hub_route_table" "pike" {
  resource_group_name = "pike"
  name                = "pike"
  virtual_hub_name    = "pike"
}
