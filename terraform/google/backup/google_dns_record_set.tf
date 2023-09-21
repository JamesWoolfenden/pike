resource "google_dns_record_set" "pike" {
  name         = "backend.freebeer.site."
  managed_zone = "prod-zone"
  type         = "A"
  ttl          = 300

  rrdatas = ["8.8.8.8"]
}

output "dns_record_set" {
  value = google_dns_record_set.pike
}
