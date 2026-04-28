resource "azurerm_monitor_data_collection_rule_association" "pike_gen" {
  name                    = "example1-dcra"
  target_resource_id      = "pike"
  data_collection_rule_id = "pike"
  description             = "example"
}
