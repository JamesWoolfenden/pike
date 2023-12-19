data "azurerm_automation_variable_int" "pike" {
  automation_account_name = "pike"
  resource_group_name     = "pike"
  name                    = "pike"
}
