data "google_cloud_run_v2_service_iam_policy" "pike" {
  name = "pike"
}

output "policy3" {
  value = data.google_cloud_run_v2_service_iam_policy.pike
}
