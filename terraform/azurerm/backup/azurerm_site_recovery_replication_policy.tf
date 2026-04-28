resource "azurerm_site_recovery_replication_policy" "pike_gen" {
  name                                                 = "policy"
  resource_group_name                                  = "pike"
  recovery_vault_name                                  = "pike"
  recovery_point_retention_in_minutes                  = 24 * 60
  application_consistent_snapshot_frequency_in_minutes = 4 * 60
}
