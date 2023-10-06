data "google_kms_crypto_key_version" "pike" {
  crypto_key = "projects/pike-gcp/locations/europe-west2/keyRings/pike/cryptoKeys/pike"
}

output "policy" {
  value = data.google_kms_crypto_key_version.pike
}
