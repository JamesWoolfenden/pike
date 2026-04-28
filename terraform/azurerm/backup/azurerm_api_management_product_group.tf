resource "azurerm_api_management_product_group" "pike_gen" {
  product_id          = "pike"
  group_name          = "pike"
  api_management_name = "pike"
  resource_group_name = "pike"
}
