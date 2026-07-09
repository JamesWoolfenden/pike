resource "google_compute_region_instant_snapshot_iam_member" "pike" {
  name   = google_compute_region_instant_snapshot.pike.name
  region = google_compute_region_instant_snapshot.pike.region

  role   = "roles/viewer"
  member = "user:james.woolfenden@gmail.com"
}
