data "google_kms_crypto_key_iam_policy" "pike" {
  crypto_key_id = "pike-gcp/europe-west2/pike/pike"
}
