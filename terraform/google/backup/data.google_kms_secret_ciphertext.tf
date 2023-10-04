data "google_kms_secret_ciphertext" "pike" {
  plaintext  = "pike"
  crypto_key = "abcdefg"
}

#deprecated
