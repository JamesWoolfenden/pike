resource "google_netapp_backup_vault" "pike" {
  name        = "pike-backup-vault"
  location    = "europe-west2"
  description = "Pike test backup vault - updated"
}
