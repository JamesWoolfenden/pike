resource "google_compute_region_url_map" "pike" {
  region          = "us-central1"
  name            = "login"
  default_service = google_compute_region_backend_service.pike.name
}
