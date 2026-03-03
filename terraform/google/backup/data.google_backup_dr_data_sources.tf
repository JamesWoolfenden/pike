data "google_backup_dr_data_sources" "pike" {
  location        = "us-central1"
  backup_vault_id = "pike"
}

output "google_backup_dr_data_sources" {
  value = data.google_backup_dr_data_sources.pike
}
