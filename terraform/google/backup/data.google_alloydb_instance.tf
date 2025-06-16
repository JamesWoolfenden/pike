data "google_alloydb_instance" "pike" {
  instance_id = "pike"
  location    = "us-central1"
  cluster_id  = "pike"
}

output "google_alloydb_instance" {
  value = data.google_alloydb_instance.pike
}
