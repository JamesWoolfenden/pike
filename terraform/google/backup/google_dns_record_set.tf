resource "google_dns_record_set" "pike" {
  name = "frontend.${google_dns_managed_zone.pike.dns_name}"
  type = "A"
  ttl  = 300

  managed_zone = google_dns_managed_zone.pike.name
  rrdatas      = ["8.8.8.8"]
}
