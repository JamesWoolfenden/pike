# Regional Secret Manager Secret
resource "google_secret_manager_regional_secret" "pike" {
  secret_id = "pike-regional-secret"
  location  = "us-central1"
  project   = "pike-477416"
}
