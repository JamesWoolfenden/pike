resource "azurerm_api_management_backend" "pike_gen" {
  name                = "example-backend"
  resource_group_name = "pike"
  api_management_name = "pike"
  protocol            = "http"
  url                 = "https://backend.com/api"
}
