data "azurerm_eventgrid_system_topic" "pike_gen" {
  name                = "eventgrid-system-topic"
  resource_group_name = "example-resources"
}
