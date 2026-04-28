resource "azurerm_windows_function_app_slot" "pike_gen" {
  name                 = "example-slot"
  function_app_id      = "pike"
  storage_account_name = "pike"

  site_config {}
}
