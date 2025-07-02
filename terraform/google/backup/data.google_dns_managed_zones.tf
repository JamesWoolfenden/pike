data "google_dns_managed_zones" "pike" {
}

output "google_dns_managed_zones" {
  value = data.google_dns_managed_zones.pike
}
