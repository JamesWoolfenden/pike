resource "google_storage_default_object_access_control" "pike" {
  bucket = "pike-example"
  role   = "READER"
  entity = "user:crwoolfenden@gmail.com"
}
