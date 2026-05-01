data "google_cloud_run_v2_service" "pike" {
  provider = google-beta
  name     = "pike"
  project  = "pike"
  location = "europe-west2"
}

output "google_cloud_run_v2_service" {
  value = data.google_cloud_run_v2_service.pike
}
