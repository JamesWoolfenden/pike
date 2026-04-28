resource "azurerm_api_management_certificate" "pike_gen" {
  name                = "example-cert"
  api_management_name = "pike"
  resource_group_name = "pike"
  data                = filebase64("example.pfx")
}
