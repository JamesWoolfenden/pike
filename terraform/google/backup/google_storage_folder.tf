# Storage Folder
resource "google_storage_bucket" "pike_folder_bucket" {
  name                        = "pike-folder-test"
  location                    = "US"
  project                     = "pike-477416"
  uniform_bucket_level_access = true
  hierarchical_namespace {
    enabled = true
  }
}

resource "google_storage_folder" "pike" {
  bucket = google_storage_bucket.pike_folder_bucket.name
  name   = "test-folder/"
}
