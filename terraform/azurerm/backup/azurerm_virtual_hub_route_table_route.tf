resource "azurerm_virtual_hub_route_table_route" "pike_gen" {
  route_table_id = "pike"

  name              = "example-route"
  destinations_type = "CIDR"
  destinations      = ["10.0.0.0/16"]
  next_hop_type     = "ResourceId"
  next_hop          = "pike"
}
