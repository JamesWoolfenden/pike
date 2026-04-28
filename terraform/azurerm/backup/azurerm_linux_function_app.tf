resource "azurerm_linux_function_app" "pike_gen" {
  name                = "example-linux-function-app"
  resource_group_name = "pike"
  location            = "pike"

  storage_account_name = "pike"
  service_plan_id      = "pike"

  site_config {}
}
