data "google_compute_region_ssl_certificate" "pike" {
  name = "pike"
}

output "cert" {
  value = data.google_compute_region_ssl_certificate.pike
}
