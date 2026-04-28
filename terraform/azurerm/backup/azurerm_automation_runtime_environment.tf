resource "azurerm_automation_runtime_environment" "pike_gen" {
  name                  = "powershell_environment_custom_config"
  automation_account_id = "pike"

  runtime_language = "PowerShell"
  runtime_version  = "7.2"

  location    = "pike"
  description = "example description"

  runtime_default_packages = {
    "az"        = "11.2.0"
    "azure cli" = "2.56.0"
  }

  tags = {
    key = "foo"
  }
}
