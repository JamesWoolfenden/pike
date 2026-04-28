resource "azurerm_automation_variable_int" "pike_gen" {
  name                    = "tfex-example-var"
  resource_group_name     = "pike"
  automation_account_name = "pike"
  value                   = 1234
}
