resource "google_secret_manager_regional_secret_iam_member" "pike" {
  project   = google_secret_manager_regional_secret.pike.project
  location  = google_secret_manager_regional_secret.pike.location
  secret_id = google_secret_manager_regional_secret.pike.secret_id
  role      = "roles/secretmanager.secretAccessor"
  member    = "user:james.woolfenden@gmail.com"
}
