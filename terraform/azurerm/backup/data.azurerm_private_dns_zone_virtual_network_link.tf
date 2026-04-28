data "azurerm_private_dns_zone_virtual_network_link" "pike_gen" {
  name                  = "test"
  resource_group_name   = "test-rg"
  private_dns_zone_name = "test-zone"
}
