data "google_oracle_database_cloud_vm_cluster" "pike" {
  location            = "us-central1"
  cloud_vm_cluster_id = "pike"
}

output "google_oracle_database_cloud_vm_cluster" {
  value = data.google_oracle_database_cloud_vm_cluster.pike
}
