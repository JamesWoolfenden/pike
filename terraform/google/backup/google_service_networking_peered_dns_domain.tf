resource "google_service_networking_peered_dns_domain" "pike" {
  name       = "pike"
  network    = google_compute_network.pike.name
  service    = "servicenetworking.googleapis.com"
  dns_suffix = "pike.example.com."

  depends_on = [google_service_networking_connection.pike]
}
