resource "azurerm_active_directory_domain_service_replica_set" "pike_gen" {
  domain_service_id = "pike"
  location          = "pike"
  subnet_id         = "pike"

  depends_on = [
    azurerm_subnet_network_security_group_association.replica,
    azurerm_virtual_network_peering.primary_replica,
    azurerm_virtual_network_peering.replica_primary,
  ]
}
