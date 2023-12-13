data "google_vmwareengine_private_cloud" "pike" {
  provider = google-beta
  name     = "pike"
  location = "us-central1"
}
