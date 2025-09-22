data "google_backup_dr_data_source_references" "pike" {
  location      = "us-central1"
  resource_type = "sqladmin.googleapis.com/Instance"
}

output "google_backup_dr_data_source_references" {
  value = data.google_backup_dr_data_source_references.pike
}
