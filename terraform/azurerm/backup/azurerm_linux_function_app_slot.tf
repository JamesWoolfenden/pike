resource "azurerm_linux_function_app_slot" "pike_gen" {
  name                 = "example-linux-function-app-slot"
  function_app_id      = "pike"
  storage_account_name = "pike"

  site_config {}
}
