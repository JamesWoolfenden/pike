data "google_compute_region_ssl_policy" "pike" {
  name = "pike"
}

output "google_compute_region_ssl_policy" {
  value = data.google_compute_region_ssl_policy.pike
}
