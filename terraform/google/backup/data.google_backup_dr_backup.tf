data "google_backup_dr_backup" "pike" {
  location        = "us-central1"
  project         = "pike-412922"
  data_source_id  = "pike"
  backup_vault_id = "pike"
}

output "google_backup_dr_backup" {
  value = data.google_backup_dr_backup.pike
}
