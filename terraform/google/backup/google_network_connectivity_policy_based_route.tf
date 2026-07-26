resource "google_network_connectivity_policy_based_route" "pike" {
  name    = "pike"
  network = google_compute_network.pike.id

  filter {
    protocol_version = "IPV4"
  }

  next_hop_other_routes = "DEFAULT_ROUTING"
}
