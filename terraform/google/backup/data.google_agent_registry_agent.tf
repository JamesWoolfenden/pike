data "google_agent_registry_agent" "pike" {
  location = "us-central1"
  agent_id = "pike"
}

output "google_agent_registry_agent" {
  value = data.google_agent_registry_agent.pike
}
