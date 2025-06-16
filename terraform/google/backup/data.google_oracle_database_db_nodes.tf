data "google_oracle_database_db_nodes" "pike" {
  cloud_vm_cluster = "pike"
  location         = "us-central1"
}

output "google_oracle_database_db_nodes" {
  value = data.google_oracle_database_db_nodes.pike
}
