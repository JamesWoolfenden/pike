data "google_iap_app_engine_service_iam_policy" "pike" {
  app_id  = "pike"
  service = "pike"
}

output "google_iap_app_engine_service_iam_policy" {
  value = data.google_iap_app_engine_service_iam_policy.pike
}
