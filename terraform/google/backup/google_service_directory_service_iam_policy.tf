data "google_iam_policy" "admin" {
  provider = google-beta
  binding {
    role = "roles/viewer"
    members = [
      "user:crwoolfenden@gmail.com",
    ]
  }
}

resource "google_service_directory_service_iam_policy" "policy" {
  provider    = google-beta
  name        = google_service_directory_service.pike.name
  policy_data = data.google_iam_policy.admin.policy_data
}
