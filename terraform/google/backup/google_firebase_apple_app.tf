resource "google_firebase_apple_app" "pike" {
  provider     = google-beta
  project      = google_firebase_project.pike.project
  display_name = "Display Name Basic"
  bundle_id    = "apple.app.12345"
  depends_on = [
    google_firebase_project.pike
  ]
}
