resource "azurerm_automation_credential" "pike_gen" {
  name                    = "credential1"
  resource_group_name     = "pike"
  automation_account_name = "pike"
  username                = "example_user"
  description             = "This is an example credential"
}
