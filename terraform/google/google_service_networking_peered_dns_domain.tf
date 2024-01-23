resource "google_service_networking_peered_dns_domain" "pike" {
  project    = 10000000
  name       = "example-com"
  network    = "default"
  dns_suffix = "example.com."
  service    = "peering-service"
}
