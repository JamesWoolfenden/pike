resource "azurerm_eventhub_namespace_schema_group" "pike_gen" {
  name                 = "example-schemaGroup"
  namespace_id         = "pike"
  schema_compatibility = "Forward"
  schema_type          = "Avro"
}
