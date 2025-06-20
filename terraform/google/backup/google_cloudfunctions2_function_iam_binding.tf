data "google_iam_policy" "admin3" {
  binding {
    role = "roles/viewer"
    members = [
      "user:james.woolfenden@gmail.com",
    ]
  }
}

resource "google_cloudfunctions2_function_iam_policy" "policy" {
  project        = google_cloudfunctions2_function.function.project
  location       = google_cloudfunctions2_function.function.location
  cloud_function = google_cloudfunctions2_function.function.name
  policy_data    = data.google_iam_policy.admin3.policy_data
}
