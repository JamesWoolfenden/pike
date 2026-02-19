resource "google_compute_health_check" "pike_health" {
  name = "pike-tcp-health-check"

  tcp_health_check {
    port = 80
  }
}

resource "google_compute_backend_service" "pike_backend" {
  name          = "pike-tcp-backend"
  protocol      = "TCP"
  health_checks = [google_compute_health_check.pike_health.id]
}

resource "google_compute_target_tcp_proxy" "pike" {
  name            = "pike-tcp-proxy"
  backend_service = google_compute_backend_service.pike_backend.id
}
