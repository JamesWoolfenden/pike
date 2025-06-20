resource "google_secret_manager_regional_secret_version" "pike" {
  secret      = google_secret_manager_regional_secret.pike.id
  secret_data = "secret-data"
}
