resource "azurerm_cdn_frontdoor_profile" "pike_gen" {
  name                     = "example-cdn-profile"
  resource_group_name      = "pike"
  sku_name                 = "Premium_AzureFrontDoor"
  response_timeout_seconds = 120

  identity {
    type         = "SystemAssigned, UserAssigned"
    identity_ids = ["pike"]
  }

  log_scrubbing_rule {
    match_variable = "RequestIPAddress"
  }

  tags = {
    environment = "Production"
  }
}
