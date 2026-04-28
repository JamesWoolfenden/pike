resource "azurerm_site_recovery_replication_recovery_plan" "pike_gen" {
  name                      = "example-recover-plan"
  recovery_vault_id         = "pike"
  source_recovery_fabric_id = "pike"
  target_recovery_fabric_id = "pike"

  shutdown_recovery_group {}

  failover_recovery_group {}

  boot_recovery_group {
    replicated_protected_items = ["pike"]
  }

}
