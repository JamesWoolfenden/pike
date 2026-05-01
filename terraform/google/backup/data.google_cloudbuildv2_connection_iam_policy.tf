data "google_cloudbuildv2_connection_iam_policy" "pike" {
  name = "pike"
}

output "google_cloudbuildv2_connection_iam_policy" {
  value = data.google_cloudbuildv2_connection_iam_policy.pike
}
