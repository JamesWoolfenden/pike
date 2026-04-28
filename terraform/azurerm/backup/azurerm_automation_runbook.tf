resource "azurerm_automation_runbook" "pike_gen" {
  name                    = "Get-AzureVMTutorial"
  location                = "pike"
  resource_group_name     = "pike"
  automation_account_name = "pike"
  log_verbose             = "true"
  log_progress            = "true"
  description             = "This is an example runbook"
  runbook_type            = "PowerShellWorkflow"

  publish_content_link {
    uri = "https://raw.githubusercontent.com/Azure/azure-quickstart-templates/c4935ffb69246a6058eb24f54640f53f69d3ac9f/101-automation-runbook-getvms/Runbooks/Get-AzureVMTutorial.ps1"
  }
}
