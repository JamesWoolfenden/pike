resource "google_compute_region_instant_snapshot_iam_binding" "pike" {
  name   = google_compute_region_instant_snapshot.pike.name
  region = google_compute_region_instant_snapshot.pike.region

  role = "roles/viewer"
  members = [
    "user:james.woolfenden@gmail.com",
  ]
}
