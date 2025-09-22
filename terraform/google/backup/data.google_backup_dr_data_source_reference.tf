data "google_backup_dr_data_source_reference" "pike" {
  data_source_reference_id = "pike"
  location                 = "us-central1"
}

output "google_backup_dr_data_source_reference" {
  value = data.google_backup_dr_data_source_reference.pike
}
