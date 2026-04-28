resource "azurerm_api_management_gateway" "pike_gen" {
  name              = "example-gateway"
  api_management_id = "pike"
  description       = "Example API Management gateway"

  location_data {
    name     = "example name"
    city     = "example city"
    district = "example district"
    region   = "example region"
  }
}
