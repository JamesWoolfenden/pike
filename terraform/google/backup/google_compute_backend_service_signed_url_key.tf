resource "google_compute_health_check" "pike_backend_health" {
  name = "pike-backend-health-check"

  tcp_health_check {
    port = 80
  }
}

resource "google_compute_backend_service" "pike_backend_svc" {
  name          = "pike-backend-service"
  health_checks = [google_compute_health_check.pike_backend_health.id]
  enable_cdn    = true
}

resource "google_compute_backend_service_signed_url_key" "pike" {
  name            = "pike-signed-url-key"
  key_value       = "MTIzNDU2Nzg5MDEyMzQ1Ng=="
  backend_service = google_compute_backend_service.pike_backend_svc.name
}
