data "google_cloud_run_locations" "pike" {}

output "google_cloud_run_locations" {
  value = data.google_cloud_run_locations.pike
}
