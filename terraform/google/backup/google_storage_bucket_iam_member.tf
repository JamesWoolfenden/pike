resource "google_storage_bucket_iam_member" "pike" {
  bucket = "pike-example"
  role   = "roles/storage.admin"
  member = "user:crwoolfenden@gmail.com"
}
