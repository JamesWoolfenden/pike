resource "azurerm_automation_source_control" "pike_gen" {
  name                  = "example"
  automation_account_id = "pike"
  folder_path           = "runbook"

  security {
    token      = "ghp_xxx"
    token_type = "PersonalAccessToken"
  }
  repository_url      = "https://github.com/foo/bat.git"
  source_control_type = "GitHub"
  branch              = "main"
}
