resource "azurerm_automation_hybrid_runbook_worker_group" "pike_gen" {
  name                    = "example"
  resource_group_name     = "pike"
  automation_account_name = "pike"
}
