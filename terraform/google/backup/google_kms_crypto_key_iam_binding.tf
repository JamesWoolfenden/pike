resource "google_kms_crypto_key_iam_binding" "pike" {
  crypto_key_id = "projects/pike-gcp/locations/europe-west2/keyRings/pike/cryptoKeys/pike"
  role          = "roles/cloudkms.cryptoKeyEncrypter"

  members = [
    "user:james.woolfenden@gmail.com",
  ]
}
