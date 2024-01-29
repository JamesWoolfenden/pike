resource "google_firebase_database_instance" "pike" {
  provider    = google-beta
  project     = google_firebase_project.pike.project
  region      = "europe-west1"
  instance_id = "active-db-jgw3"
  depends_on = [
    google_firebase_project.pike
  ]
}
