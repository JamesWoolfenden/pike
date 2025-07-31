data "google_dns_managed_zones" "pike" {
  # project="pike-412922"
}

output "google_dns_managed_zones" {
  value = data.google_dns_managed_zones.pike
}
