resource "azurerm_api_management_group_user" "pike_gen" {
  user_id             = "pike"
  group_name          = "example-group"
  resource_group_name = "pike"
  api_management_name = "pike"
}
