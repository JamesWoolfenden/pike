resource "google_service_directory_namespace_iam_binding" "binding2" {
  provider = google-beta
  name     = google_service_directory_namespace.example.name
  role     = "roles/viewer"
  members = [
    "user:crwoolfenden@gmail.com",
  ]
}
