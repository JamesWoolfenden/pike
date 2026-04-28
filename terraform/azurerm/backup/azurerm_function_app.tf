resource "azurerm_function_app" "pike_gen" {
  name                 = "test-azure-functions"
  location             = "pike"
  resource_group_name  = "pike"
  app_service_plan_id  = "pike"
  storage_account_name = "pike"
}
