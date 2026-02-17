resource "tls_private_key" "pike_ssl" {
  algorithm = "RSA"
  rsa_bits  = 2048
}

resource "tls_self_signed_cert" "pike_ssl" {
  private_key_pem = tls_private_key.pike_ssl.private_key_pem

  subject {
    common_name  = "pike-test.example.com"
    organization = "Pike Testing"
  }

  validity_period_hours = 8760 # 1 year

  allowed_uses = [
    "key_encipherment",
    "digital_signature",
    "server_auth",
  ]
}

resource "google_compute_ssl_certificate" "pike" {
  name        = "pike-ssl-cert"
  private_key = tls_private_key.pike_ssl.private_key_pem
  certificate = tls_self_signed_cert.pike_ssl.cert_pem
}
