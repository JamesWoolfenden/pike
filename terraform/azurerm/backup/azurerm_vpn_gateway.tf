resource "azurerm_vpn_gateway" "pike_gen" {
  name                = "example-vpng"
  location            = "pike"
  resource_group_name = "pike"
  virtual_hub_id      = "pike"
}
