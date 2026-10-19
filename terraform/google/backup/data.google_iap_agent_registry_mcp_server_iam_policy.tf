data "google_iap_agent_registry_mcp_server_iam_policy" "pike" {
  mcp_server_id = "pike-test-mcp-server"
}

output "google_iap_agent_registry_mcp_server_iam_policy" {
  value = data.google_iap_agent_registry_mcp_server_iam_policy.pike
}