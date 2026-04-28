data "azurerm_cognitive_account_project" "pike_gen" {
  name                   = "example-project"
  cognitive_account_name = "example-account"
  resource_group_name    = "example-resources"
}
