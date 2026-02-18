resource "google_compute_managed_ssl_certificate" "pike" {
  name = "pike-managed-ssl-cert"

  managed {
    domains = ["pike-test.example.com"]
  }
}
