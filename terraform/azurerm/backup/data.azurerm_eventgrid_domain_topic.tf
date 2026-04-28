data "azurerm_eventgrid_domain_topic" "pike_gen" {
  name                = "my-eventgrid-domain-topic"
  resource_group_name = "example-resources"
}
