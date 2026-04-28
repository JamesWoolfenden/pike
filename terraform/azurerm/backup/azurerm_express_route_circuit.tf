resource "azurerm_express_route_circuit" "pike_gen" {
  name                  = "expressRoute1"
  resource_group_name   = "pike"
  location              = "pike"
  service_provider_name = "Equinix"
  peering_location      = "Silicon Valley"
  bandwidth_in_mbps     = 50

  sku {
    tier   = "Standard"
    family = "MeteredData"
  }

  tags = {
    environment = "Production"
  }
}
