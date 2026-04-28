resource "azurerm_windows_web_app" "pike_gen" {
  name                = "example"
  resource_group_name = "pike"
  location            = "pike"
  service_plan_id     = "pike"

  site_config {}
}
