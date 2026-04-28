resource "azurerm_frontdoor_custom_https_configuration" "pike_gen" {
  frontend_endpoint_id              = "pike" ["exampleFrontendEndpoint1"]
  custom_https_provisioning_enabled = false
}
