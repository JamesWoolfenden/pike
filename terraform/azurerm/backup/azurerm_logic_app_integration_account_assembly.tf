resource "azurerm_logic_app_integration_account_assembly" "pike_gen" {
  name                     = "example-assembly"
  resource_group_name      = "pike"
  integration_account_name = "pike"
  assembly_name            = "TestAssembly"
  content                  = filebase64("testdata/log4net.dll")
}
