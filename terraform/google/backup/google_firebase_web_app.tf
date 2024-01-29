resource "google_firebase_web_app" "pike" {
  provider     = google-beta
  project      = google_firebase_project.pike.project
  display_name = "Display Name Basic"
  depends_on = [
  google_firebase_project.pike]
}
