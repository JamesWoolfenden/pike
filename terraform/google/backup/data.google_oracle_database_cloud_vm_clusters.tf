data "google_oracle_database_cloud_vm_clusters" "pike" {
  location = "us-central1"
}

output "google_oracle_database_cloud_vm_clusters" {
  value = data.google_oracle_database_cloud_vm_clusters.pike
}
