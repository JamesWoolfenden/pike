resource "azurerm_traffic_manager_azure_endpoint" "pike_gen" {
  name                 = "example-endpoint"
  profile_id           = "pike"
  always_serve_enabled = true
  weight               = 100
  target_resource_id   = "pike"
}
