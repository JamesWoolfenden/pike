data "google_compute_region_target_http_proxy" "pike" {
  name   = "pike"
  region = "us-central1"
}

output "google_compute_region_target_http_proxy" {
  value = data.google_compute_region_target_http_proxy.pike
}
