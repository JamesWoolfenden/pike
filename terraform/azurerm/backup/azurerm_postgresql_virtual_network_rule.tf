resource "azurerm_postgresql_virtual_network_rule" "pike_gen" {
  name                                 = "postgresql-vnet-rule"
  resource_group_name                  = "pike"
  server_name                          = "pike"
  subnet_id                            = "pike"
  ignore_missing_vnet_service_endpoint = true
}
