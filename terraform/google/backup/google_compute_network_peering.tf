resource "google_compute_network" "pike_network_1" {
  name                    = "pike-network-1"
  auto_create_subnetworks = false
}

resource "google_compute_network" "pike_network_2" {
  name                    = "pike-network-2"
  auto_create_subnetworks = false
}

resource "google_compute_network_peering" "pike" {
  name         = "pike-peering-1-to-2"
  network      = google_compute_network.pike_network_1.id
  peer_network = google_compute_network.pike_network_2.id
}
