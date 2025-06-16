data "google_kms_crypto_key_latest_version" "pike" {
  crypto_key = "pike/us-central1/pike/pike"
}

output "google_kms_crypto_key_latest_version" {
  value = data.google_kms_crypto_key_latest_version.pike
}
