resource "google_filestore_snapshot" "pike" {
  name     = "test-snapshot"
  instance = google_filestore_instance.pike.name
  location = "us-central1"
}
