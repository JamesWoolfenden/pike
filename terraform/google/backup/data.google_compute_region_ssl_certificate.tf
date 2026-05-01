data "google_compute_region_ssl_certificate" "pike" {
  name = "pike"
}

output "google_compute_region_ssl_certificate" {
  value = data.google_compute_region_ssl_certificate.pike
}
