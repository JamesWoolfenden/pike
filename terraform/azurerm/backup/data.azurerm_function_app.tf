data "azurerm_function_app" "pike_gen" {
  name                = "test-azure-functions"
  resource_group_name = "pike"
}
