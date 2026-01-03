data "google_compute_region_security_policy" "pike" {
  name = "pike"
}

output "google_compute_region_security_policy" {
  value = data.google_compute_region_security_policy.pike
}
