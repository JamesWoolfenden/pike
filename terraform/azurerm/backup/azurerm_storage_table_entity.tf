resource "azurerm_storage_table_entity" "pike_gen" {
  storage_table_id = "pike"

  partition_key = "examplepartition"
  row_key       = "examplerow"

  entity = {
    example = "example"
  }
}
