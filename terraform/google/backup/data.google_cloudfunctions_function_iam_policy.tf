data "google_cloudfunctions_function_iam_policy" "pike" {
  cloud_function = "pike"
}

output "policy" {
  value = data.google_cloudfunctions_function_iam_policy.pike
}
