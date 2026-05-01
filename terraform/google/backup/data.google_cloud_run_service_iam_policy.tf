data "google_cloud_run_service_iam_policy" "pike" {
  service = "pike"
}

output "google_cloud_run_service_iam_policy" {
  value = data.google_cloud_run_service_iam_policy.pike
}
