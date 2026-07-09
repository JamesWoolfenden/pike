data "google_iam_policy" "google_compute_region_instant_snapshot" {
  binding {
    role = "roles/viewer"
    members = [
      "user:james.woolfenden@gmail.com",
    ]
  }
}

resource "google_compute_region_instant_snapshot_iam_policy" "pike" {
  name   = google_compute_region_instant_snapshot.pike.name
  region = google_compute_region_instant_snapshot.pike.region

  policy_data = data.google_iam_policy.google_compute_region_instant_snapshot.policy_data
}
