data "google_kms_crypto_key_versions" "pike" {
  crypto_key = "pike/us-central1/pike/pike"
}

output "google_kms_crypto_key_versions" {
  value = data.google_kms_crypto_key_versions.pike
}
