resource "azurerm_virtual_hub" "pike_gen" {
  name                = "example-virtualhub"
  resource_group_name = "pike"
  location            = "pike"
  virtual_wan_id      = "pike"
  address_prefix      = "10.0.0.0/23"
}
