data "azurerm_automation_runbook" "pike_gen" {
  name                    = "existing-runbook"
  resource_group_name     = "existing"
  automation_account_name = "existing-automation"
}
