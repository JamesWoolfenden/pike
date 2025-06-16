data "google_oracle_database_cloud_exadata_infrastructures" "pike" {
  location = "us-central1"
}

output "google_oracle_database_cloud_exadata_infrastructures" {
  value = data.google_oracle_database_cloud_exadata_infrastructures.pike
}
