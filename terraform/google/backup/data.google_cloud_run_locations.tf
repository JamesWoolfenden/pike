data "google_cloud_run_locations" "pike" {}

output "locations" {
  value = data.google_cloud_run_locations.pike
}
