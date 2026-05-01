data "google_dns_managed_zone" "pike" {
  name = "prod-zone"
}

output "google_dns_managed_zone" {
  value = data.google_dns_managed_zone.pike
}
