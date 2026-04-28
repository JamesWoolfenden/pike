resource "azurerm_data_protection_backup_instance_data_lake_storage" "pike_gen" {
  name                               = "example-data-protection-backup-instance-data-lake-storage"
  data_protection_backup_vault_id    = "pike"
  location                           = "pike"
  storage_account_id                 = "pike"
  backup_policy_data_lake_storage_id = "pike"
  storage_container_names            = ["pike", "pike"]

  depends_on = ["pike"]
}
