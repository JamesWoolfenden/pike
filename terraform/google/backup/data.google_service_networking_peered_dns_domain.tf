data "google_service_networking_peered_dns_domain" "pike" {
  service = "pike"
  project = "pike-477416"
  network = "pike"
  name    = "pike"
}

output "google_service_networking_peered_dns_domain" {
  value = data.google_service_networking_peered_dns_domain.pike
}
