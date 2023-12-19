data "azurerm_automation_variable_bool" "pike" {
  automation_account_name = "pike"
  resource_group_name     = "pike"
  name                    = "pike"
}
