resource "azurerm_api_management_api_version_set" "pike_gen" {
  name                = "example-apimapi-1_0_0"
  resource_group_name = "pike"
  api_management_name = "pike"
  display_name        = "ExampleAPIVersionSet"
  versioning_scheme   = "Segment"
}
