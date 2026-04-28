resource "azurerm_express_route_port_authorization" "pike_gen" {
  name                    = "exampleERCAuth"
  express_route_port_name = "pike"
  resource_group_name     = "pike"
}
