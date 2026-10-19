data "google_cloudbuild_worker_pool" "pike" {
  name     = "pike-test-pool"
  location = "europe-west2"
}

output "google_cloudbuild_worker_pool" {
  value = data.google_cloudbuild_worker_pool.pike
}