data "google_iam_policy" "admin" {
  binding {
    role = "roles/viewer"
    members = [
      "user:james.woolfenden@gmail.com",
    ]
  }
}

resource "google_cloudfunctions_function_iam_policy" "policy" {
  project        = google_cloudfunctions_function.pikey.project
  region         = google_cloudfunctions_function.pikey.region
  cloud_function = google_cloudfunctions_function.pikey.name
  policy_data    = data.google_iam_policy.admin.policy_data
}
