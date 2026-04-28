resource "azurerm_backup_protected_vm" "pike_gen" {
  resource_group_name = "pike"
  recovery_vault_name = "pike"
  source_vm_id        = "pike"
  backup_policy_id    = "pike"
}
