resource "azurerm_automation_webhook" "pike_gen" {
  name                    = "TestRunbook_webhook"
  resource_group_name     = "pike"
  automation_account_name = "pike"
  expiry_time             = "2021-12-31T00:00:00Z"
  enabled                 = true
  runbook_name            = "pike"
  parameters = {
    input = "parameter"
  }
}
