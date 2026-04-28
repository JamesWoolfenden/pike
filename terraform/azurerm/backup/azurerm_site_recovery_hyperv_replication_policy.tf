resource "azurerm_site_recovery_hyperv_replication_policy" "pike_gen" {
  name                                               = "policy"
  recovery_vault_id                                  = "pike"
  recovery_point_retention_in_hours                  = 2
  application_consistent_snapshot_frequency_in_hours = 1
  replication_interval_in_seconds                    = 300
}
