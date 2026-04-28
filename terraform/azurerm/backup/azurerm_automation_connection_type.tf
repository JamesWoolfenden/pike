resource "azurerm_automation_connection_type" "pike_gen" {
  name                    = "example"
  resource_group_name     = "pike"
  automation_account_name = "pike"

  field {
    name = "example"
    type = "string"
  }
}
