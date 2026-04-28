resource "azurerm_api_management_workspace_api_version_set" "pike_gen" {
  name                        = "example-version-set"
  api_management_workspace_id = "pike"
  display_name                = "Example API Version Set"
  versioning_scheme           = "Segment"
}
