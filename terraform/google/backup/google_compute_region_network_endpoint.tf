# Region Network Endpoint - endpoint for regional load balancing
resource "google_compute_network" "pike_network" {
  name                    = "pike-test-network"
  auto_create_subnetworks = true
  project                 = "pike-477416"
}

resource "google_compute_region_network_endpoint_group" "pike_neg" {
  name                  = "pike-region-neg"
  region                = "us-central1"
  project               = "pike-477416"
  network_endpoint_type = "INTERNET_IP_PORT"
  network               = google_compute_network.pike_network.id
}

resource "google_compute_region_network_endpoint" "pike" {
  region                        = "us-central1"
  project                       = "pike-477416"
  region_network_endpoint_group = google_compute_region_network_endpoint_group.pike_neg.name

  port       = 443
  ip_address = "8.8.8.8"
}
