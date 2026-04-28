resource "azurerm_api_management_subscription" "pike_gen" {
  api_management_name = "pike"
  resource_group_name = "pike"
  user_id             = "pike"
  product_id          = "pike"
  display_name        = "Parser API"
}
