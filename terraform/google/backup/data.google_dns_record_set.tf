data "google_dns_record_set" "pike" {
  name         = "pike.freebeer.site."
  managed_zone = "prod-zone"
  type         = "A"
}

output "google_dns_record_set" {
  value = data.google_dns_record_set.pike
}
