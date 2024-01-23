data "google_iam_policy" "admin2" {
  provider = google-beta
  binding {
    role = "roles/viewer"
    members = [
      "user:crwoolfenden@gmail.com",
    ]
  }
}

resource "google_service_directory_namespace_iam_policy" "policy" {
  provider    = google-beta
  name        = google_service_directory_namespace.example.name
  policy_data = data.google_iam_policy.admin2.policy_data
}
