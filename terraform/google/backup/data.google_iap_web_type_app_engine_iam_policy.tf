data "google_iap_web_type_app_engine_iam_policy" "pike" {
  app_id = "asdasd"
}

output "policy2" {
  value = data.google_iap_web_type_app_engine_iam_policy.pike
}
