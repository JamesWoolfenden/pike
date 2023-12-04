data "google_tpu_v2_runtime_versions" "pike" {
  provider = google-beta
  zone     = "us-central1-b"
}
