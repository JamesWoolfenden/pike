data "google_compute_storage_pool" "pike" {
  zone = "us-central1-a"
  name = "pike"
}

output "google_compute_storage_pool" {
  value = data.google_compute_storage_pool.pike
}
