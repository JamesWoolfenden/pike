data "google_kms_crypto_key" "pike" {
  name     = "pike-key-keep"
  key_ring = "projects/pike-361314/locations/europe-west1/keyRings/pike"
}

output "pike-key" {
  value = data.google_kms_crypto_key.pike
}
