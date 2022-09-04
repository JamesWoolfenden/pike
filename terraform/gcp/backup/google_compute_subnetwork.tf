resource "google_compute_subnetwork" "pike" {
  description   = "test for pike"
  name          = "sub-pike"
  network       = "default"
  ip_cidr_range = "10.0.1.0/24"
  log_config {
    metadata = "INCLUDE_ALL_METADATA"
  }
}
