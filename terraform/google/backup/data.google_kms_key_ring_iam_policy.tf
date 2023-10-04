data "google_kms_key_ring_iam_policy" "pike" {
  key_ring_id = "projects/pike-gcp/locations/europe-west2/keyRings/pike"
}

output "keyring" {
  value = data.google_kms_key_ring_iam_policy.pike
}
