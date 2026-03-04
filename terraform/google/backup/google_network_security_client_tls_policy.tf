resource "google_network_security_client_tls_policy" "pike" {
  name        = "pike-client-tls-policy"
  location    = "europe-west2"
  description = "Pike test client TLS policy - updated"
}
