resource "google_compute_health_check" "pike_ssl_health" {
  name = "pike-ssl-health-check"

  tcp_health_check {
    port = 443
  }
}

resource "google_compute_backend_service" "pike_ssl_backend" {
  name          = "pike-ssl-backend"
  protocol      = "SSL"
  health_checks = [google_compute_health_check.pike_ssl_health.id]
}

resource "google_compute_ssl_certificate" "pike_cert" {
  name        = "pike-test-cert"
  private_key = file("${path.module}/test-key.pem")
  certificate = file("${path.module}/test-cert.pem")
}

resource "google_compute_target_ssl_proxy" "pike" {
  name             = "pike-ssl-proxy"
  backend_service  = google_compute_backend_service.pike_ssl_backend.id
  ssl_certificates = [google_compute_ssl_certificate.pike_cert.id]
}
