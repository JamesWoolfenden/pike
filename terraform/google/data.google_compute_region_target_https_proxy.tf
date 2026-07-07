data "google_compute_region_target_https_proxy" "pike" {
  name   = "pike"
  region = "us-central1"
}

output "google_compute_region_target_https_proxy" {
  value = data.google_compute_region_target_https_proxy.pike
}
