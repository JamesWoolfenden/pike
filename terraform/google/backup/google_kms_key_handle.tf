resource "google_kms_key_handle" "pike" {
  provider               = google-beta
  project                = "pike-412922"
  name                   = "tf-test-key-handle"
  location               = "global"
  resource_type_selector = "storage.googleapis.com/Bucket"
}
