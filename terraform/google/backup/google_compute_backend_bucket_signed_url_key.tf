resource "google_storage_bucket" "pike_backend_bucket" {
  name     = "pike-backend-bucket-${random_id.suffix.hex}"
  location = "US"
}

resource "random_id" "suffix" {
  byte_length = 4
}

resource "google_compute_backend_bucket" "pike_backend" {
  name        = "pike-backend-bucket"
  bucket_name = google_storage_bucket.pike_backend_bucket.name
  enable_cdn  = true
}

resource "google_compute_backend_bucket_signed_url_key" "pike" {
  name           = "pike-signed-url-key"
  key_value      = "MTIzNDU2Nzg5MDEyMzQ1Ng=="
  backend_bucket = google_compute_backend_bucket.pike_backend.name
}
