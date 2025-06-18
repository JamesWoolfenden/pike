resource "google_cloud_run_domain_mapping" "pike" {
  location = "us-central1"
  name     = "verified-domain.com"

  metadata {
    namespace = "pike-412922"
  }

  spec {
    route_name = google_cloud_run_service.pike.name
  }
}
