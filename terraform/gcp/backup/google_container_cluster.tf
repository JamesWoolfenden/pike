resource "google_container_cluster" "pike" {
  name               = "pike"
  initial_node_count = 1
}
