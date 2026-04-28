resource "azurerm_kusto_script" "pike_gen" {
  name                               = "example"
  database_id                        = "pike"
  url                                = "pike"
  continue_on_errors_enabled         = true
  force_an_update_when_value_changed = "first"
  script_level                       = "Database"
  principal_permissions_action       = "RemovePermissionOnScriptCompletion"
}
