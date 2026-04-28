resource "azurerm_function_app_slot" "pike_gen" {
  name                 = "test-azure-functions_slot"
  location             = "pike"
  resource_group_name  = "pike"
  app_service_plan_id  = "pike"
  function_app_name    = "pike"
  storage_account_name = "pike"
}
