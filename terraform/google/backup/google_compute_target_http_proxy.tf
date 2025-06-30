resource "google_compute_target_http_proxy" "pike" {
  name    = "test-proxy"
  url_map = google_compute_url_map.urlmap.id
}
