# resource "google_compute_network" "net" {
#   name                    = "my-network"
#   auto_create_subnetworks = false
# }

# resource "google_compute_subnetwork" "subnet" {
#   name          = "my-subnetwork"
#   network       = google_compute_network.net.id
#   ip_cidr_range = "10.0.0.0/16"
#   region        = "us-central1"
# }

# resource "google_compute_router" "router" {
#   name    = "my-router"
#   region  = google_compute_subnetwork.subnet.region
#   network = google_compute_network.net.id
# }

resource "google_compute_router_route_policy" "rp-export" {
  router = google_compute_router.router.name
  region = google_compute_router.router.region
  name   = "my-rp1"
  type   = "ROUTE_POLICY_TYPE_EXPORT"
  terms {
    priority = 1
    match {
      expression = "destination == '10.0.0.0/12'"
    }
    actions {
      expression = "accept()"
    }
  }
}
