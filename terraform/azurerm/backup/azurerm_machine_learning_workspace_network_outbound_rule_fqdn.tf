resource "azurerm_machine_learning_workspace_network_outbound_rule_fqdn" "pike_gen" {
  name             = "example-outboundrule"
  workspace_id     = "pike"
  destination_fqdn = "example.com"
}
