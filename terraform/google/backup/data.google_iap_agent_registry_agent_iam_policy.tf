data "google_iap_agent_registry_agent_iam_policy" "pike" {
  agent_id = "pike-test-agent"
}

output "google_iap_agent_registry_agent_iam_policy" {
  value = data.google_iap_agent_registry_agent_iam_policy.pike
}