data "azurerm_automation_variable_string" "pike" {
  automation_account_name = "pike"
  resource_group_name     = "pike"
  name                    = "pike"
}
