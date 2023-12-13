data "google_compute_backend_bucket_iam_policy" "pike" {
  provider = google-beta
  name     = "pike"
}
