data "google_compute_storage_pool_types" "pike" {
  storage_pool_type = "pike"
  zone              = "us-central1-a"
}

output "google_compute_storage_pool_types" {
  value = data.google_compute_storage_pool_types.pike
}
