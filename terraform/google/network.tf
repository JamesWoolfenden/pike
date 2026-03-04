# Reusable VPC network for testing network-dependent resources
resource "google_compute_network" "pike_test" {
  name                    = "pike-test-network"
  auto_create_subnetworks = false
  description             = "Pike test network for resource testing"
}

resource "google_compute_subnetwork" "pike_test" {
  name          = "pike-test-subnet"
  ip_cidr_range = "10.0.0.0/24"
  region        = "europe-west2"
  network       = google_compute_network.pike_test.id
}
