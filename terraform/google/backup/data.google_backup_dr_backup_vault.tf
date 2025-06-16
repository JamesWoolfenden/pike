data "google_backup_dr_backup_vault" "pike" {
  location        = "us-central1"
  backup_vault_id = "pike"
}

output "google_backup_dr_backup_vault" {
  value = data.google_backup_dr_backup_vault.pike
}
