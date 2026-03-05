resource "random_id" "cache_suffix" {
  byte_length = 4
}

resource "google_storage_bucket" "cache_bucket" {
  name                        = "pike-cache-bucket-${random_id.cache_suffix.hex}"
  location                    = "US"
  uniform_bucket_level_access = true
}

resource "google_storage_anywhere_cache" "cache" {
  bucket            = google_storage_bucket.cache_bucket.name
  anywhere_cache_id = "pike-cache"
  zone              = "us-central1-f"
  ttl               = "3601s"
  admission_policy  = "admit-on-second-miss"
}
