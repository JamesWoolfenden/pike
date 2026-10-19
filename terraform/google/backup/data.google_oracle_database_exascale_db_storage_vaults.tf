data "google_oracle_database_exascale_db_storage_vaults" "pike" {
  location = "europe-west2"
}

output "google_oracle_database_exascale_db_storage_vaults" {
  value = data.google_oracle_database_exascale_db_storage_vaults.pike
}