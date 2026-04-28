resource "azurerm_api_management_workspace_named_value" "pike_gen" {
  name                        = "example-named-value"
  api_management_workspace_id = "pike"
  display_name                = "ExampleProperty"
  value                       = "Example Value"
  tags                        = ["tag1", "tag2", "tag3"]
}
