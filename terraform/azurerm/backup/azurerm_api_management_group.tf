resource "azurerm_api_management_group" "pike_gen" {
  name                = "example-apimg"
  resource_group_name = "pike"
  api_management_name = "pike"
  display_name        = "Example Group"
  description         = "This is an example API management group."
}
