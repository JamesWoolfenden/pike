resource "azurerm_automation_job_schedule" "pike_gen" {
  resource_group_name     = "tf-rgr-automation"
  automation_account_name = "tf-automation-account"
  schedule_name           = "hour"
  runbook_name            = "Get-VirtualMachine"

  parameters = {
    resourcegroup = "tf-rgr-vm"
    vmname        = "TF-VM-01"
  }
}
