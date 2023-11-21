resource "google_storage_object_access_control" "pike" {
  object = "results.txt"
  bucket = "pike-example"
  role   = "READER"
  entity = "allUsers"
}
