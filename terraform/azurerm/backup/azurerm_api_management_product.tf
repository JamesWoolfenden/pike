resource "azurerm_api_management_product" "pike_gen" {
  product_id            = "test-product"
  api_management_name   = "pike"
  resource_group_name   = "pike"
  display_name          = "Test Product"
  subscription_required = true
  approval_required     = true
  published             = true
}
