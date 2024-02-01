resource "google_cloud_run_v2_job_iam_member" "pike" {

  name   = google_cloud_run_v2_job.default.name
  role   = "roles/viewer"
  member = "user:crwoolfenden@gmail.com"
}
