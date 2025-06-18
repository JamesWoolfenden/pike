resource "google_cloud_run_v2_worker_pool_iam_binding" "pike" {
  project  = google_cloud_run_v2_worker_pool.pike.project
  location = google_cloud_run_v2_worker_pool.pike.location
  name     = google_cloud_run_v2_worker_pool.pike.name
  role     = "roles/viewer"
  members = [
    "user:james.woolfenden@gmail.com",
  ]
}
