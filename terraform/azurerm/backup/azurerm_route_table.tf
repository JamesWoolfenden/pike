resource "azurerm_route_table" "pike_gen" {
  name                = "example-route-table"
  location            = "pike"
  resource_group_name = "pike"

  route {
    name           = "route1"
    address_prefix = "10.1.0.0/16"
    next_hop_type  = "VnetLocal"
  }

  tags = {
    environment = "Production"
  }
}
