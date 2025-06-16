data "google_kms_ekm_connection_iam_policy" "pike" {
  name = "projects/pike/locations/us-central1/ekmConnections/pike"
}

output "google_kms_ekm_connection_iam_policy" {
  value = data.google_kms_ekm_connection_iam_policy.pike
}
