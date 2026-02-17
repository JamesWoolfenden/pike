resource "google_compute_node_template" "pike" {
  name      = "pike-node-template-test"
  region    = "us-central1"
  node_type = "n1-node-96-624"
}
