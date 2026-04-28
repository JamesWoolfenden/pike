data "azurerm_vpn_server_configuration" "pike_gen" {
  name                = "existing-local-vpn-server-configuration"
  resource_group_name = "existing-resource-group"
}
