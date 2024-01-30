resource "google_api_gateway_gateway" "pike" {
  provider   = google-beta
  api_config = google_api_gateway_api_config.pike.id
  gateway_id = "pike"
}
