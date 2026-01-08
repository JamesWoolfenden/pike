resource "google_compute_router_peer" "pike" {
  name                      = "my-router-peer"
  router                    = google_compute_router.router.name
  region                    = "us-central1"
  peer_asn                  = 65513
  advertised_route_priority = 100
  interface                 = "interface-1"
}
