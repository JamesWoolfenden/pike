resource "google_storage_bucket_access_control" "pike" {
  bucket = "pike-example"
  role   = "READER"
  entity = "user:crwoolfenden@gmail.com"
}
