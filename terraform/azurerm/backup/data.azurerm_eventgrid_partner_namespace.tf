data "azurerm_eventgrid_partner_namespace" "pike_gen" {
  name                = "my-eventgrid-partner-namespace"
  resource_group_name = "example-resources"
}
