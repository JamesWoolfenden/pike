data "google_iam_policy" "admin" {
  binding {
    role = "roles/cloudkms.cryptoKeyEncrypter"

    members = [
      "user:james.woolfenden@gmail.com",
    ]
  }
}

resource "google_kms_crypto_key_iam_policy" "crypto_key" {
  crypto_key_id = "projects/pike-gcp/locations/europe-west2/keyRings/pike/cryptoKeys/pike"
  policy_data   = data.google_iam_policy.admin.policy_data
}
