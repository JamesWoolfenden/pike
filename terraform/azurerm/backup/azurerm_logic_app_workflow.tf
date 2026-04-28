resource "azurerm_logic_app_workflow" "pike_gen" {
  name                = "workflow1"
  location            = "pike"
  resource_group_name = "pike"
}
