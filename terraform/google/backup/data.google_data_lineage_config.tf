data "google_data_lineage_config" "pike" {
  parent   = "projects/pike-gcp"
  location = "us-central1"
}

output "google_data_lineage_config" {
  value = data.google_data_lineage_config.pike
}
