data "azurerm_eventgrid_domain" "pike_gen" {
  name                = "my-eventgrid-domain"
  resource_group_name = "example-resources"
}
