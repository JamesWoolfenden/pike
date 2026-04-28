resource "azurerm_api_management_api_schema" "pike_gen" {
  api_name            = "pike"
  api_management_name = "pike"
  resource_group_name = "pike"
  schema_id           = "example-schema"
  content_type        = "application/vnd.ms-azure-apim.xsd+xml"
  value               = file("api_management_api_schema.xml")
}
