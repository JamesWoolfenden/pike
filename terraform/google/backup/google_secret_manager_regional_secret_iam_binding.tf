resource "google_secret_manager_regional_secret_iam_binding" "pike" {
  project   = google_secret_manager_regional_secret.pike.project
  secret_id = google_secret_manager_regional_secret.pike.secret_id
  role      = "roles/secretmanager.secretAccessor"
  members = [
    "user:james.woolfenden@gmail.com",
  ]
  depends_on = [
    google_secret_manager_regional_secret.pike
  ]
}
