resource "azurerm_route" "pike_gen" {
  name                = "acceptanceTestRoute1"
  resource_group_name = "pike"
  route_table_name    = "pike"
  address_prefix      = "10.1.0.0/16"
  next_hop_type       = "VnetLocal"
}
