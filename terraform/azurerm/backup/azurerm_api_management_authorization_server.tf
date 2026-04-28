resource "azurerm_api_management_authorization_server" "pike_gen" {
  name                         = "test-server"
  api_management_name          = "pike"
  resource_group_name          = "pike"
  display_name                 = "Test Server"
  authorization_endpoint       = "https://example.mydomain.com/client/authorize"
  client_id                    = "42424242-4242-4242-4242-424242424242"
  client_registration_endpoint = "https://example.mydomain.com/client/register"

  grant_types = [
    "authorizationCode",
  ]
  authorization_methods = [
    "GET",
  ]
}
