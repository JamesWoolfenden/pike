resource "azurerm_network_manager_routing_rule_collection" "pike_gen" {
  name                     = "example-routing-rule-collection"
  routing_configuration_id = "pike"
  network_group_ids        = ["azurerm_network_manager_network_group.example.id"]
  description              = "example routing rule collection"
}
