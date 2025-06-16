data "google_clouddeploy_custom_target_type_iam_policy" "pike" {
  name = "pike"
}

output "google_clouddeploy_custom_target_type_iam_policy" {
  value = data.google_clouddeploy_custom_target_type_iam_policy.pike
}
