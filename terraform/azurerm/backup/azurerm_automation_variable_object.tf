resource "azurerm_automation_variable_object" "pike_gen" {
  name                    = "tfex-example-var"
  resource_group_name     = "pike"
  automation_account_name = "pike"
  value = jsonencode({
    greeting = "Hello, Terraform Basic Test."
    language = "en"
  })
}
