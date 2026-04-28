resource "azurerm_app_configuration_feature" "pike_gen" {
  configuration_store_id = "pike"
  description            = "test description"
  name                   = "test-ackey"
  label                  = "test-ackeylabel"
  enabled                = true
}
