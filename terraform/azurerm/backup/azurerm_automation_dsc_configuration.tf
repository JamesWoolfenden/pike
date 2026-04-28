resource "azurerm_automation_dsc_configuration" "pike_gen" {
  name                    = "test"
  resource_group_name     = "pike"
  automation_account_name = "pike"
  location                = "pike"
  content_embedded        = "configuration test {}"
}
