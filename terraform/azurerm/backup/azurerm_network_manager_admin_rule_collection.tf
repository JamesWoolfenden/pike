resource "azurerm_network_manager_admin_rule_collection" "pike_gen" {
  name                            = "example-admin-rule-collection"
  security_admin_configuration_id = "pike"
  network_group_ids               = ["pike"]
}
