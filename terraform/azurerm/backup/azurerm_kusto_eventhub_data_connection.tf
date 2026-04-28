resource "azurerm_kusto_eventhub_data_connection" "pike_gen" {
  name                = "my-kusto-eventhub-data-connection"
  resource_group_name = "pike"
  location            = "pike"
  cluster_name        = "pike"
  database_name       = "pike"

  eventhub_id    = "pike"
  consumer_group = "pike"

  table_name           = "my-table"         #(Optional)
  mapping_rule_name    = "my-table-mapping" #(Optional)
  data_format          = "JSON"             #(Optional)
  retrieval_start_date = "2023-06-26T12:00:00Z"
}
