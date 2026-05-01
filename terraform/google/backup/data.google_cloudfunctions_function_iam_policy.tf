data "google_cloudfunctions_function_iam_policy" "pike" {
  cloud_function = "pike"
}

output "google_cloudfunctions_function_iam_policy" {
  value = data.google_cloudfunctions_function_iam_policy.pike
}
