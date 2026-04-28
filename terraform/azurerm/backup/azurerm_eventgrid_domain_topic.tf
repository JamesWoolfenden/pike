resource "azurerm_eventgrid_domain_topic" "pike_gen" {
  name                = "my-eventgrid-domain-topic"
  domain_name         = "pike"
  resource_group_name = "pike"
}
