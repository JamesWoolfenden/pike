data "google_oracle_database_db_servers" "pike" {
  cloud_exadata_infrastructure = "pike"
  location                     = "us-central1"
}

output "google_oracle_database_db_servers" {
  value = data.google_oracle_database_db_servers.pike
}
