resource "google_compute_network" "pike_vpn_network" {
  name                    = "pike-vpn-network"
  auto_create_subnetworks = false
}

resource "google_compute_ha_vpn_gateway" "pike" {
  name    = "pike-ha-vpn-gateway"
  region  = "us-central1"
  network = google_compute_network.pike_vpn_network.id
}
