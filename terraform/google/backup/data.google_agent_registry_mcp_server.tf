data "google_agent_registry_mcp_server" "pike" {
  location      = "us-central1"
  mcp_server_id = "pike"
}

output "google_agent_registry_mcp_server" {
  value = data.google_agent_registry_mcp_server.pike
}
