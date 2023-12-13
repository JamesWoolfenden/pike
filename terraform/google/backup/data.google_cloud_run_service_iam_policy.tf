data "google_cloud_run_service_iam_policy" "pike" {
  service = "pike"
}

output "policy" {
  value = data.google_cloud_run_service_iam_policy.pike
}
