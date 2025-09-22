data "google_iap_web_forwarding_rule_service_iam_policy" "pike" {
  forwarding_rule_service_name = "pike"
}

output "google_iap_web_forwarding_rule_service_iam_policy" {
  value = data.google_iap_web_forwarding_rule_service_iam_policy.pike
}
