resource "azurerm_dev_center_network_connection" "pike_gen" {
  name                = "example-dcnc"
  resource_group_name = "pike"
  location            = "pike"
  domain_join_type    = "AzureADJoin"
  subnet_id           = "pike"
}
