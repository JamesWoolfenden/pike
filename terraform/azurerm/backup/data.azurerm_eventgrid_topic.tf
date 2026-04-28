data "azurerm_eventgrid_topic" "pike_gen" {
  name                = "my-eventgrid-topic"
  resource_group_name = "example-resources"
}
