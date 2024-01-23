resource "google_service_directory_service_iam_binding" "pike" {
  provider = google-beta
  name     = google_service_directory_service.pike.name
  role     = "roles/viewer"
  members = [
    "user:crwoolfenden@gmail.com",
  ]
}
