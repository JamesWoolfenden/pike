resource "tls_private_key" "rsa" {
  algorithm = "RSA"
}

# Public key loaded from a terraform-generated private key, using the PEM (RFC 1421) format
data "tls_public_key" "pike" {
  private_key_pem = tls_private_key.rsa.private_key_pem
}
