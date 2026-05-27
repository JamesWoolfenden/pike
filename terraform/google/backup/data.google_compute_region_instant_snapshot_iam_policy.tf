data "google_compute_region_instant_snapshot_iam_policy" "pike" {
  name   = "pike"
  region = "us-central1"
}

output "google_compute_region_instant_snapshot_iam_policy" {
  value = data.google_compute_region_instant_snapshot_iam_policy.pike
}
