resource "google_storage_bucket" "bucket" {
  name     = "example-storage-bucket-jgw"
  location = "US"
}

data "google_site_verification_token" "token" {
  type                = "SITE"
  identifier          = "https://${google_storage_bucket.bucket.name}.storage.googleapis.com/"
  verification_method = "FILE"
}

resource "google_storage_bucket_object" "object" {
  name    = data.google_site_verification_token.token.token
  content = "google-site-verification: ${data.google_site_verification_token.token.token}"
  bucket  = google_storage_bucket.bucket.name
}

resource "google_storage_object_access_control" "public_rule" {
  bucket = google_storage_bucket.bucket.name
  object = google_storage_bucket_object.object.name
  role   = "READER"
  entity = "allUsers"
}

resource "google_site_verification_web_resource" "example" {
  site {
    type       = data.google_site_verification_token.token.type
    identifier = data.google_site_verification_token.token.identifier
  }
  verification_method = data.google_site_verification_token.token.verification_method
}

resource "google_site_verification_owner" "example" {
  web_resource_id = google_site_verification_web_resource.example.id
  email           = "user@example.com"
}
