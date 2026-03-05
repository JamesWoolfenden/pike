resource "google_storage_bucket" "pike_managed_folder_test" {
  name                        = "pike-managed-folder-test-${random_id.bucket_suffix.hex}"
  location                    = "US"
  uniform_bucket_level_access = true
}

resource "random_id" "bucket_suffix" {
  byte_length = 4
}

resource "google_storage_managed_folder" "pike" {
  bucket = google_storage_bucket.pike_managed_folder_test.name
  name   = "pike/pike/"
}
