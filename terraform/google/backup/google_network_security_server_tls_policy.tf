resource "google_network_security_server_tls_policy" "pike" {
  name     = "pike-server-tls-policy"
  location = "europe-west2"
}
