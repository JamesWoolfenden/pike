data "google_cloud_run_v2_service_iam_policy" "pike" {
  name = "pike"
}

output "google_cloud_run_v2_service_iam_policy" {
  value = data.google_cloud_run_v2_service_iam_policy.pike
}
