resource "azurerm_automation_hybrid_runbook_worker" "pike_gen" {
  resource_group_name     = "pike"
  automation_account_name = "pike"
  worker_group_name       = "pike"
  vm_resource_id          = "pike"
  worker_id               = "00000000-0000-0000-0000-000000000000" #unique uuid
}
