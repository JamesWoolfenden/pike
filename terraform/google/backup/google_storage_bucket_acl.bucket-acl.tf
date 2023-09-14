resource "google_storage_bucket_acl" "bucketacl" {
  bucket = "test-bucket-jgw-today"

  predefined_acl = "projectprivate"
}
