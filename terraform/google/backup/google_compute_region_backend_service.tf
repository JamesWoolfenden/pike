resource "google_compute_region_backend_service" "pike" {
  region                = "us-central1"
  name                  = "region-service"
  health_checks         = [google_compute_region_health_check.pike.id]
  protocol              = "HTTPS"
  load_balancing_scheme = "INTERNAL_MANAGED"
  locality_lb_policy    = "ROUND_ROBIN"
}
