resource "azurerm_virtual_hub_route_table" "pike_gen" {
  name           = "example-vhubroutetable"
  virtual_hub_id = "pike"
  labels         = ["label1"]

  route {
    name              = "example-route"
    destinations_type = "CIDR"
    destinations      = ["10.0.0.0/16"]
    next_hop_type     = "ResourceId"
    next_hop          = "pike"
  }
}
