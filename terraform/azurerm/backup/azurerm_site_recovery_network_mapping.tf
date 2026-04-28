resource "azurerm_site_recovery_network_mapping" "pike_gen" {
  name                        = "recovery-network-mapping-1"
  resource_group_name         = "pike"
  recovery_vault_name         = "pike"
  source_recovery_fabric_name = "primary-fabric"
  target_recovery_fabric_name = "secondary-fabric"
  source_network_id           = "pike"
  target_network_id           = "pike"
}
