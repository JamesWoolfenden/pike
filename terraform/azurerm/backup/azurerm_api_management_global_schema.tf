resource "azurerm_api_management_global_schema" "pike_gen" {
  schema_id           = "example-schema1"
  api_management_name = "pike"
  resource_group_name = "pike"
  type                = "xml"
  value               = file("api_management_api_schema.xml")
}
