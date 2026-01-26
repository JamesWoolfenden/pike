resource "google_app_engine_service_network_settings" "pike" {
  service = "pike"
  network_settings {
    ingress_traffic_allowed = "INGRESS_TRAFFIC_ALLOWED_INTERNAL_ONLY"
  }
}
