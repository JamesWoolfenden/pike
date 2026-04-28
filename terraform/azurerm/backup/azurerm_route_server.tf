resource "azurerm_route_server" "pike_gen" {
  name                             = "example-routerserver"
  resource_group_name              = "pike"
  location                         = "pike"
  sku                              = "Standard"
  public_ip_address_id             = "pike"
  subnet_id                        = "pike"
  branch_to_branch_traffic_enabled = true
  hub_routing_preference           = "ASPath"
}
