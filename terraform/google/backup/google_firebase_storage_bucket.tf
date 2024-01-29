resource "google_firebase_storage_bucket" "pike" {
  provider  = google-beta
  project   = google_firebase_project.pike.project
  bucket_id = google_storage_bucket.default.id
  depends_on = [
    google_storage_bucket.default
  ]
}

resource "google_storage_bucket" "default" {
  provider                    = google-beta
  name                        = "test_bucket_jgw_firebase"
  location                    = "EU"
  uniform_bucket_level_access = true
}
