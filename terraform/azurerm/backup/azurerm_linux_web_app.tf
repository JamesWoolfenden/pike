resource "azurerm_linux_web_app" "pike_gen" {
  name                = "example"
  resource_group_name = "pike"
  location            = "pike"
  service_plan_id     = "pike"

  site_config {}
}
