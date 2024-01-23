resource "google_service_directory_service_iam_member" "pike" {
  provider = google-beta
  name     = google_service_directory_service.pike.name
  role     = "roles/viewer"
  member   = "user:crwoolfenden@gmail.com"
}
