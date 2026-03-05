data "google_backup_dr_management_server" "pike" {
  location = "us-central1"
}

output "google_backup_dr_management_server" {
  value = data.google_backup_dr_management_server.pike
}
