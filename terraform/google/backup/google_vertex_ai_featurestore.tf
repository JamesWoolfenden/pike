resource "google_vertex_ai_featurestore" "pike" {
  name = "terraform"
  labels = {
    foo = "bar"
  }
  region = "us-central1"
  online_serving_config {
    fixed_node_count = 2
  }
  encryption_spec {
    kms_key_name = "projects/pike-gcp/locations/us-central1/keyRings/cluster-1vpUU/cryptoKeys/cluster-1vpUU"
  }
  force_destroy = true
}
