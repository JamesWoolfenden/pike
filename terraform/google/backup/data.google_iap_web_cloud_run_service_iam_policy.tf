data "google_iap_web_cloud_run_service_iam_policy" "pike" {
  cloud_run_service_name = "pike"
}

output "google_iap_web_cloud_run_service_iam_policy" {
  value = data.google_iap_web_cloud_run_service_iam_policy.pike
}
