resource "google_vertex_ai_metadata_store" "store" {
  provider    = google-beta
  name        = "test-store"
  description = "Store to test the terraform module"
  region      = "us-central1"
}
