resource "google_compute_network" "pike" {
  name                            = "pike"
  description                     = "pike network"
  auto_create_subnetworks         = true
  mtu                             = 1460
  routing_mode                    = "REGIONAL"
  delete_default_routes_on_create = false
}
