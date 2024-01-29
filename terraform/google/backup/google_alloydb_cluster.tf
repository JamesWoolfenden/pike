resource "google_alloydb_cluster" "default" {
  cluster_id = "alloydb-cluster"
  location   = "europe-west2"
  network    = google_compute_network.default.id

  initial_user {
    password = "cluster_secret"
  }
}
