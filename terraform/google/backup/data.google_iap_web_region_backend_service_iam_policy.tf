data "google_iap_web_region_backend_service_iam_policy" "pike" {
  web_region_backend_service = "pike"
}

output "policy2" {
  value = data.google_iap_web_region_backend_service_iam_policy.pike
}
