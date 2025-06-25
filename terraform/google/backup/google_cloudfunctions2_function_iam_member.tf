resource "google_cloudfunctions2_function_iam_member" "member" {
  project        = google_cloudfunctions2_function.function.project
  location       = google_cloudfunctions2_function.function.location
  cloud_function = google_cloudfunctions2_function.function.name
  role           = "roles/viewer"
  member         = "user:james.woolfenden@gmail.com"
}
