data "google_dns_managed_zones" "pike" {
  # project="pike-477416"
}

output "google_dns_managed_zones" {
  value = data.google_dns_managed_zones.pike
}
