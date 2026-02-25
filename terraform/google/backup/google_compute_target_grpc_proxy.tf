resource "google_compute_health_check" "pike_grpc_health" {
  name = "pike-grpc-health-check"

  grpc_health_check {
    port = 443
  }
}

resource "google_compute_backend_service" "pike_grpc_backend" {
  name                  = "pike-grpc-backend"
  protocol              = "GRPC"
  load_balancing_scheme = "INTERNAL_SELF_MANAGED"
  health_checks         = [google_compute_health_check.pike_grpc_health.id]
}

resource "google_compute_url_map" "pike_grpc" {
  name            = "pike-grpc-url-map"
  default_service = google_compute_backend_service.pike_grpc_backend.id
}

resource "google_compute_target_grpc_proxy" "pike" {
  name    = "pike-grpc-proxy"
  url_map = google_compute_url_map.pike_grpc.id
}
