resource "azurerm_kusto_eventgrid_data_connection" "pike_gen" {
  name                         = "my-kusto-eventgrid-data-connection"
  resource_group_name          = "pike"
  location                     = "pike"
  cluster_name                 = "pike"
  database_name                = "pike"
  storage_account_id           = "pike"
  eventhub_id                  = "pike"
  eventhub_consumer_group_name = "pike"

  table_name        = "my-table"         #(Optional)
  mapping_rule_name = "my-table-mapping" #(Optional)
  data_format       = "JSON"             #(Optional)

  depends_on = ["pike"]
}
