resource "azurerm_mssql_managed_instance_failover_group" "pike_gen" {
  name                        = "example-failover-group"
  location                    = "pike"
  managed_instance_id         = "pike"
  partner_managed_instance_id = "pike"
  secondary_type              = "Geo"
  read_write_endpoint_failover_policy {
    mode          = "Automatic"
    grace_minutes = 60
  }
  depends_on = [
    azurerm_private_dns_zone_virtual_network_link.primary,
    azurerm_private_dns_zone_virtual_network_link.failover,
  ]
}
