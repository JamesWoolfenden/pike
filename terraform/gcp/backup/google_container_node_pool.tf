resource "google_container_node_pool" "pike" {
  cluster = google_container_cluster.pike.name
}
