resource "google_storage_bucket_iam_binding" "binding" {
  bucket = "test-bucket-jgw-today"
  role   = "roles/storage.admin"
  members = [
    "user:james.woolfenden@gmail.com",
    "serviceAccount:pike-service@pike-361314.iam.gserviceaccount.com"
  ]

}
