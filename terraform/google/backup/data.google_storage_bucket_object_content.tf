data "google_storage_bucket_object_content" "pike" {
  bucket = "pike-example"
  name   = "results.txt"
}
