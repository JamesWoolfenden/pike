resource "azurerm_site_recovery_protection_container_mapping" "pike_gen" {
  name                                      = "container-mapping"
  resource_group_name                       = "pike"
  recovery_vault_name                       = "pike"
  recovery_fabric_name                      = "pike"
  recovery_source_protection_container_name = "pike"
  recovery_target_protection_container_id   = "pike"
  recovery_replication_policy_id            = "pike"
}
