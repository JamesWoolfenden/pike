data "google_iap_web_type_app_engine_iam_policy" "pike" {
  app_id = "asdasd"
}

output "google_iap_web_type_app_engine_iam_policy" {
  value = data.google_iap_web_type_app_engine_iam_policy.pike
}
