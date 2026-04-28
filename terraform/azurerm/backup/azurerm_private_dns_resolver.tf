resource "azurerm_private_dns_resolver" "pike_gen" {
  name                = "example"
  resource_group_name = "pike"
  location            = "pike"
  virtual_network_id  = "pike"
}
