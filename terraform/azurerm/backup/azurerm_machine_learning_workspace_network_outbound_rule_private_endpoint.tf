resource "azurerm_machine_learning_workspace_network_outbound_rule_private_endpoint" "pike_gen" {
  name                = "example-outboundrule"
  workspace_id        = "pike"
  service_resource_id = "pike"
  sub_resource_target = "blob"
}
