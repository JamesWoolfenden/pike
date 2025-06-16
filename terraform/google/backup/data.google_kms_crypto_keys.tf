data "google_kms_crypto_keys" "pike" {
  key_ring = "us-central1/pike"
}

output "google_kms_crypto_keys" {
  value = data.google_kms_crypto_keys.pike
}
