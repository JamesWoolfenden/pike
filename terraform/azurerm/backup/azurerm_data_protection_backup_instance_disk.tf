resource "azurerm_data_protection_backup_instance_disk" "pike_gen" {
  name                         = "example-backup-instance"
  location                     = "pike"
  vault_id                     = "pike"
  disk_id                      = "pike"
  snapshot_resource_group_name = "pike"
  backup_policy_id             = "pike"
}
