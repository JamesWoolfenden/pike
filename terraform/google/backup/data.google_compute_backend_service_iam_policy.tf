data "google_compute_backend_service_iam_policy" "pike" {
  provider = google-beta
  name     = "pike"
}
