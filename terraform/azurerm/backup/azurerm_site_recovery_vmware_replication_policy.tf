resource "azurerm_site_recovery_vmware_replication_policy" "pike_gen" {
  name                                                 = "example-policy"
  recovery_vault_id                                    = "pike"
  recovery_point_retention_in_minutes                  = 1440
  application_consistent_snapshot_frequency_in_minutes = 240
}
