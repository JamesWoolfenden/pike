resource "azurerm_app_configuration_key" "pike_gen" {
  configuration_store_id = "pike"
  key                    = "appConfKey1"
  label                  = "somelabel"
  value                  = "a test"

  depends_on = [
    azurerm_role_assignment.appconf_dataowner
  ]
}
