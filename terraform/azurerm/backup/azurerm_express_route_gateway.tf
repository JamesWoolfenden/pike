resource "azurerm_express_route_gateway" "pike_gen" {
  name                = "expressRoute1"
  resource_group_name = "pike"
  location            = "pike"
  virtual_hub_id      = "pike"
  scale_units         = 1

  tags = {
    environment = "Production"
  }
}
