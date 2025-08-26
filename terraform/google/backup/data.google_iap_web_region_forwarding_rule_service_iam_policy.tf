data "google_iap_web_region_forwarding_rule_service_iam_policy" "pike" {
  provider                   = google-beta
  web_region_backend_service = "pike"
}

output "google_iap_web_region_forwarding_rule_service_iam_policy" {
  value = data.google_iap_web_region_forwarding_rule_service_iam_policy.pike
}
