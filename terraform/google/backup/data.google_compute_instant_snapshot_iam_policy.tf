data "google_compute_instant_snapshot_iam_policy" "pike" {
  name = "pike"
  zone = "us-central1-a"
}

output "google_compute_instant_snapshot_iam_policy" {
  value = data.google_compute_instant_snapshot_iam_policy.pike
}
