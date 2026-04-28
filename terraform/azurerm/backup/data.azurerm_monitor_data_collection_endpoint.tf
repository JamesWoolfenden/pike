data "azurerm_monitor_data_collection_endpoint" "pike_gen" {
  name                = "example-mdce"
  resource_group_name = "pike"
}
