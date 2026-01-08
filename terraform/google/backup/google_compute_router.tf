resource "google_compute_router" "router" {
  name    = "my-router"
  region  = google_compute_subnetwork.subnet.region
  network = google_compute_network.net.id

  bgp {
    asn = 64514
  }
}
