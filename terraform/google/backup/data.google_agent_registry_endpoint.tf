data "google_agent_registry_endpoint" "pike" {
  location    = "us-central1"
  endpoint_id = "pike"
}

output "google_agent_registry_endpoint" {
  value = data.google_agent_registry_endpoint.pike
}
