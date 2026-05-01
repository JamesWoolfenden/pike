data "google_kms_key_ring_iam_policy" "pike" {
  key_ring_id = "projects/pike-gcp/locations/europe-west2/keyRings/pike"
}

output "google_kms_key_ring_iam_policy" {
  value = data.google_kms_key_ring_iam_policy.pike
}
