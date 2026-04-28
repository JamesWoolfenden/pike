resource "azurerm_logic_app_standard" "pike_gen" {
  name                 = "example-logic-app"
  location             = "pike"
  resource_group_name  = "pike"
  app_service_plan_id  = "pike"
  storage_account_name = "pike"

  app_settings = {
    "FUNCTIONS_WORKER_RUNTIME"     = "node"
    "WEBSITE_NODE_DEFAULT_VERSION" = "~18"
  }
}
