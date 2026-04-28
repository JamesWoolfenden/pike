data "azurerm_site_recovery_replication_recovery_plan" "pike_gen" {
  name              = "example-recovery-plan"
  recovery_vault_id = "pike"
}
