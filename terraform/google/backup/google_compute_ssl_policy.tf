resource "google_compute_ssl_policy" "pike" {
  name            = "pike-test-ssl-policy"
  profile         = "MODERN"
  min_tls_version = "TLS_1_2"
}
