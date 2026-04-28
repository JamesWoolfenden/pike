resource "azurerm_traffic_manager_nested_endpoint" "pike_gen" {
  name                    = "example-endpoint"
  target_resource_id      = "pike"
  priority                = 1
  profile_id              = "pike"
  minimum_child_endpoints = 9
  weight                  = 5
}
