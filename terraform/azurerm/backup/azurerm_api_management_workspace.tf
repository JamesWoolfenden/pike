resource "azurerm_api_management_workspace" "pike_gen" {
  name              = "example-workspace"
  api_management_id = "pike"
  display_name      = "my workspace"
}
