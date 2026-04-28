resource "azurerm_network_manager_verifier_workspace_reachability_analysis_intent" "pike_gen" {
  name                    = "example-intent"
  verifier_workspace_id   = "pike"
  source_resource_id      = "pike"
  destination_resource_id = "pike"
  description             = "example"
  ip_traffic {
    source_ips        = ["10.0.2.1"]
    source_ports      = ["80"]
    destination_ips   = ["10.0.2.2"]
    destination_ports = ["*"]
    protocols         = ["Any"]
  }
}
