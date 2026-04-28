resource "azurerm_automation_variable_bool" "pike_gen" {
  name                    = "tfex-example-var"
  resource_group_name     = "pike"
  automation_account_name = "pike"
  value                   = false
}
