data "azurerm_logic_app_workflow" "pike_gen" {
  name                = "workflow1"
  resource_group_name = "my-resource-group"
}
