resource "azurerm_private_dns_zone_virtual_network_link" "pike_gen" {
  name                  = "test"
  resource_group_name   = "pike"
  private_dns_zone_name = "pike"
  virtual_network_id    = "pike"
}
