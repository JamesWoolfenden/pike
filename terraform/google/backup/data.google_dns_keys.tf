data "google_dns_keys" "pike" {
  //managed_zone = "private.${google_dns_managed_zone.pike.dns_name}"
  managed_zone = "prod-zone"
}

output "keys" {
  value = data.google_dns_keys.pike
}
