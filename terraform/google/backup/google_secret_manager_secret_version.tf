resource "google_secret_manager_secret_version" "pike" {
  secret_data = "mysecret"
  secret      = google_secret_manager_secret.pike.id
}
