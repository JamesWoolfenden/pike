data "google_cloud_run_v2_worker_pool" "pike" {
  name     = "pike"
  location = "us-central1"
}

output "google_cloud_run_v2_worker_pool" {
  value = data.google_cloud_run_v2_worker_pool.pike
}
