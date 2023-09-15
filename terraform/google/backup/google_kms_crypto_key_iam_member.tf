resource "google_kms_crypto_key_iam_member" "pike" {
  crypto_key_id = "projects/pike-gcp/locations/europe-west2/keyRings/pike/cryptoKeys/pike"
  role          = "roles/cloudkms.cryptoKeyEncrypter"
  member        = "user:james.woolfenden@gmail.com"
}
