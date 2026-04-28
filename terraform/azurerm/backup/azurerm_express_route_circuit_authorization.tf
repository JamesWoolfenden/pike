resource "azurerm_express_route_circuit_authorization" "pike_gen" {
  name                       = "exampleERCAuth"
  express_route_circuit_name = "pike"
  resource_group_name        = "pike"
}
