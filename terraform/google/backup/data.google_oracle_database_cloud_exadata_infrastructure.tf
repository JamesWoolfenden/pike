data "google_oracle_database_cloud_exadata_infrastructure" "pike" {
  location                        = "us-central1"
  cloud_exadata_infrastructure_id = "pike"
}

output "google_oracle_database_cloud_exadata_infrastructure" {
  value = data.google_oracle_database_cloud_exadata_infrastructure.pike
}
