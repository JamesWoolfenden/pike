resource "google_storage_bucket" "bucket" {
  name                        = "my_bucket_jgw_2025"
  location                    = "eu"
  force_destroy               = true
  uniform_bucket_level_access = true
}

resource "google_storage_bucket_object" "metadata_folder" {
  name    = "metadata/"
  content = " "
  bucket  = google_storage_bucket.bucket.name
}

resource "google_biglake_database" "database" {
  name    = "pike"
  catalog = google_biglake_catalog.catalog.id
  type    = "HIVE"
  hive_options {
    location_uri = "gs://${google_storage_bucket.bucket.name}/${google_storage_bucket_object.metadata_folder.name}"
    parameters = {
      "owner" : "James Woolfenden"
    }
  }
}
