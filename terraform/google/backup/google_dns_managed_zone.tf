resource "google_dns_managed_zone" "pike" {
  name     = "prod-zone"
  dns_name = "freebeer.site."
  labels = {
    pike = "permissions"
  }
}

output "zone" {
  value = google_dns_managed_zone.pike
}
