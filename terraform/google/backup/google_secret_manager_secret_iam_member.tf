resource "google_secret_manager_secret_iam_member" "pike" {
  member    = "user:anonymous@gmail.com"
  secret_id = google_secret_manager_secret.pike.id
  role      = "roles/secretmanager.secretAccessor"
}
