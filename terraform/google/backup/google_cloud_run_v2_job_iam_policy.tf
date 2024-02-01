data "google_iam_policy" "admin6" {
  binding {
    role = "roles/viewer"
    members = [
      "user:crwoolfenden@gmail.com",
    ]
  }
}

resource "google_cloud_run_v2_job_iam_policy" "policy" {
  name        = google_cloud_run_v2_job.default.name
  policy_data = data.google_iam_policy.admin6.policy_data
}
