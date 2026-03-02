# Storage Object ACL
resource "google_storage_bucket" "pike_acl_bucket" {
  name     = "pike-acl-bucket"
  location = "US"
  project  = "pike-477416"
}

resource "google_storage_bucket_object" "pike_object" {
  name    = "test-object.txt"
  bucket  = google_storage_bucket.pike_acl_bucket.name
  content = "Pike test content"
}

resource "google_storage_object_acl" "pike" {
  bucket = google_storage_bucket.pike_acl_bucket.name
  object = google_storage_bucket_object.pike_object.name

  role_entity = [
    "READER:allUsers",
    "OWNER:user-pike-service@pike-477416.iam.gserviceaccount.com",
  ]
}
