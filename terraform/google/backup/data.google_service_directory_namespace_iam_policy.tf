data "google_service_directory_namespace_iam_policy" "pike" {
  provider = google-beta
  name     = google_service_directory_namespace.example.name
}
