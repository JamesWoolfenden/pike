resource "azurerm_data_factory_dataset_snowflake" "pike_gen" {
  name                = "example"
  data_factory_id     = "pike"
  linked_service_name = "pike"

  schema_name = "foo_schema"
  table_name  = "foo_table"
}
