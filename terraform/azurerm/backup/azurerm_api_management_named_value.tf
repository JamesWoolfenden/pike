resource "azurerm_api_management_named_value" "pike_gen" {
  name                = "example-apimg"
  resource_group_name = "pike"
  api_management_name = "pike"
  display_name        = "ExampleProperty"
  value               = "Example Value"
}
