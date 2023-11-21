data "google_cloud_run_service" "pike" {
  location = "europe-west2"
  name     = "pike"
}

output "service" {
  value = data.google_cloud_run_service.pike
}
