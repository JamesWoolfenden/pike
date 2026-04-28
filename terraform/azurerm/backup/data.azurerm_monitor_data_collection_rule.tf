data "azurerm_monitor_data_collection_rule" "pike_gen" {
  name                = "example-rule"
  resource_group_name = "pike"
}
