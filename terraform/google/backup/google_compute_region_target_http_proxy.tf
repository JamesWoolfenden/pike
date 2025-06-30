resource "google_compute_region_target_http_proxy" "pike" {
  region  = "us-central1"
  name    = "test-proxy"
  url_map = google_compute_region_url_map.pike.id
}
