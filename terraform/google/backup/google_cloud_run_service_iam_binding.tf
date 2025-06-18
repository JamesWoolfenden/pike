resource "google_cloud_run_service_iam_binding" "pike" {
  project  = google_cloud_run_service.pike.project
  location = google_cloud_run_service.pike.location
  service  = google_cloud_run_service.pike.name
  role     = "roles/viewer"
  members = [
    "user:james.woolfenden@gmail.com",
  ]
}
