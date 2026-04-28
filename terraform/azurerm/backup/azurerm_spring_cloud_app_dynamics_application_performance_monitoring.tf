resource "azurerm_spring_cloud_app_dynamics_application_performance_monitoring" "pike_gen" {
  name                    = "example"
  spring_cloud_service_id = "pike"
  agent_account_name      = "example-agent-account-name"
  controller_host_name    = "example-controller-host-name"
  agent_application_name  = "example-agent-application-name"
  agent_tier_name         = "example-agent-tier-name"
  agent_node_name         = "example-agent-node-name"
  agent_unique_host_id    = "example-agent-unique-host-id"
  controller_ssl_enabled  = true
  controller_port         = 8080
  globally_enabled        = true
}
