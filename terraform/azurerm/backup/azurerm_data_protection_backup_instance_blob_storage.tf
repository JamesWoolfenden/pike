resource "azurerm_data_protection_backup_instance_blob_storage" "pike_gen" {
  name               = "example-backup-instance"
  vault_id           = "pike"
  location           = "pike"
  storage_account_id = "pike"
  backup_policy_id   = "pike"

  depends_on = ["pike"]
}
