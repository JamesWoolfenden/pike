resource "azurerm_machine_learning_workspace_network_outbound_rule_service_tag" "pike_gen" {
  name         = "example-outboundrule"
  workspace_id = "pike"
  service_tag  = "AppService"
  protocol     = "TCP"
  port_ranges  = "443"
}
