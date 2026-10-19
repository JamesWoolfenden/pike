data "google_iap_agent_registry_endpoint_iam_policy" "pike" {
  endpoint_id = "pike-test-endpoint"
}

output "google_iap_agent_registry_endpoint_iam_policy" {
  value = data.google_iap_agent_registry_endpoint_iam_policy.pike
}