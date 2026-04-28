resource "azurerm_eventgrid_partner_namespace" "pike_gen" {
  name                    = "example-partner-namespace"
  location                = "pike"
  resource_group_name     = "pike"
  partner_registration_id = "pike"
}
