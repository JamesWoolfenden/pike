resource "google_vertex_ai_endpoint" "pike" {
  name         = "endpoint-name"
  display_name = "sample-endpoint"
  description  = "A sample vertex endpoint"
  location     = "us-central1"
  region       = "us-central1"
  labels = {
    label-one = "value-one"
  }
  network = "projects/${data.google_project.project.number}/global/networks/${google_compute_network.vertex_network.name}"
  encryption_spec {
    kms_key_name = "projects/pike-gcp/locations/us-central1/keyRings/cluster-1vpUU/cryptoKeys/cluster-1vpUU"
  }
  depends_on = [
    google_service_networking_connection.vertex_vpc_connection
  ]
}
