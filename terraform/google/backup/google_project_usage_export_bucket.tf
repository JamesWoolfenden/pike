resource "google_project_usage_export_bucket" "pike" {
  project     = "pike-412922"
  bucket_name = google_storage_bucket.temp.name
}

resource "google_storage_bucket" "temp" {
  name                        = "usage-tracking-bucket-jgw"
  location                    = "europe-west2"
  uniform_bucket_level_access = true
}
