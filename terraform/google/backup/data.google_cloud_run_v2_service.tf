data "google_cloud_run_v2_service" "pike" {
  provider = google-beta
  name     = "pike"
  project  = "pike-gcp"
  location = "europe-west2"
}

output "service2" {
  value = data.google_cloud_run_v2_service.pike
}
