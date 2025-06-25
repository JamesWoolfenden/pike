resource "google_kms_crypto_key_version" "pike" {
  crypto_key = google_kms_crypto_key.cryptokey.id
}
