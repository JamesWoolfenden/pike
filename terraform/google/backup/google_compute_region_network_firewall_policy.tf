resource "google_compute_region_network_firewall_policy" "pike" {
  name        = "pike"
  region      = "us-central1"
  description = "Test region network firewall policy for pike"
}
