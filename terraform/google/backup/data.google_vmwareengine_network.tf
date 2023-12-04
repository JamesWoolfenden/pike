data "google_vmwareengine_network" "pike" {
  provider = google-beta
  location = "us-central1"
  name     = "pike"
}
