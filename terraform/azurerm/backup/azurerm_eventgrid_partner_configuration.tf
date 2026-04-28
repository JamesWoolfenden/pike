resource "azurerm_eventgrid_partner_configuration" "pike_gen" {
  resource_group_name                     = "pike"
  default_maximum_expiration_time_in_days = 14

  partner_authorization {
    partner_registration_id              = "804a11ca-ce9b-4158-8e94-3c8dc7a072ec"
    partner_name                         = "Auth0"
    authorization_expiration_time_in_utc = "2025-02-05T00:00:00Z"
  }

  tags = {
    environment = "Production"
  }
}
