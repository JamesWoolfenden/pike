data "google_alloydb_cluster" "pike" {
  cluster_id = "pike"
}

output "google_alloydb_cluster" {
  value = data.google_alloydb_cluster.pike
}
