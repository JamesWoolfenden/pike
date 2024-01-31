resource "google_cloud_run_v2_service_iam_binding" "pike" {

  name = google_cloud_run_v2_service.pike.name
  role = "roles/viewer"
  members = [
    "user:crwoolfenden@gmail.com",
  ]
}
