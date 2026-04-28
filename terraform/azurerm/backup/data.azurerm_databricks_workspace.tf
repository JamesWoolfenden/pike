data "azurerm_databricks_workspace" "pike_gen" {
  name                = "example-workspace"
  resource_group_name = "example-rg"
}
