resource "azurerm_digital_twins_time_series_database_connection" "pike_gen" {
  name                            = "example-connection"
  digital_twins_id                = "pike"
  eventhub_name                   = "pike"
  eventhub_namespace_id           = "pike"
  eventhub_namespace_endpoint_uri = "sb://${azurerm_eventhub_namespace.example.name}.servicebus.windows.net"
  eventhub_consumer_group_name    = "pike"
  kusto_cluster_id                = "pike"
  kusto_cluster_uri               = "pike"
  kusto_database_name             = "pike"
  kusto_table_name                = "exampleTable"

  depends_on = [
    azurerm_role_assignment.database_contributor,
    azurerm_role_assignment.eventhub_data_owner,
    azurerm_kusto_database_principal_assignment.example
  ]
}
