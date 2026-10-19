data "google_oracle_database_exascale_db_storage_vault" "pike" {
  exascale_db_storage_vault_id = "pike-test-vault1"
  location                     = "europe-west2"
}

output "google_oracle_database_exascale_db_storage_vault" {
  value = data.google_oracle_database_exascale_db_storage_vault.pike
}