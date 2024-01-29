resource "google_firebase_hosting_site" "pike" {
  provider = google-beta
  project  = google_firebase_project.pike.project
  site_id  = "pike3"
  depends_on = [
    google_firebase_project.pike
  ]
}
