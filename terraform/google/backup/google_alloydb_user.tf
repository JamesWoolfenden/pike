resource "google_alloydb_user" "user1" {
  provider  = google-beta
  cluster   = google_alloydb_cluster.default.name
  user_id   = "user1"
  user_type = "ALLOYDB_BUILT_IN"

  password       = "user_secret"
  database_roles = ["alloydbsuperuser"]
  depends_on     = [google_alloydb_instance.default]
}


data "google_project" "project" {}

resource "google_compute_network" "default" {
  name = "alloydb-network"
}

resource "google_compute_global_address" "private_ip_alloc" {
  name          = "alloydb-cluster"
  address_type  = "INTERNAL"
  purpose       = "VPC_PEERING"
  prefix_length = 16
  network       = google_compute_network.default.id
}

resource "google_service_networking_connection" "vpc_connection" {
  network                 = google_compute_network.default.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.private_ip_alloc.name]
}
