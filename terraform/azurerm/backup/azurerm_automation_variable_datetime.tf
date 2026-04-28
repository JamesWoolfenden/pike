resource "azurerm_automation_variable_datetime" "pike_gen" {
  name                    = "tfex-example-var"
  resource_group_name     = "pike"
  automation_account_name = "pike"
  value                   = "2019-04-24T21:40:54.074Z"
}
