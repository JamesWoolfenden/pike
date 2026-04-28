resource "azurerm_logic_app_integration_account_schema" "pike_gen" {
  name                     = "example-ias"
  resource_group_name      = "pike"
  integration_account_name = "pike"
  content                  = file("testdata/integration_account_schema_content.xsd")
}
