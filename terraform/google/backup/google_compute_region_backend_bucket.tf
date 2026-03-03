# Compute Region Backend Bucket
resource "google_storage_bucket" "pike_backend" {
  name     = "pike-region-backend-bucket"
  location = "europe-west2"
  project  = "pike-477416"
}

resource "google_compute_region_backend_bucket" "pike" {
  provider              = google-beta
  name                  = "pike-region-backend"
  region                = "europe-west2"
  bucket_name           = google_storage_bucket.pike_backend.name
  project               = "pike-477416"
  load_balancing_scheme = "EXTERNAL_MANAGED"
}
