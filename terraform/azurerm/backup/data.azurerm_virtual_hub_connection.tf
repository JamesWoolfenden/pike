data "azurerm_virtual_hub_connection" "pike" {
  resource_group_name = "pike"
  name                = "pike"
  virtual_hub_name    = "pike"
}
