resource "azurerm_kusto_iothub_data_connection" "pike_gen" {
  name                = "my-kusto-iothub-data-connection"
  resource_group_name = "pike"
  location            = "pike"
  cluster_name        = "pike"
  database_name       = "pike"

  iothub_id                 = "pike"
  consumer_group            = "pike"
  shared_access_policy_name = "pike"
  event_system_properties   = ["message-id", "sequence-number", "to"]

  table_name           = "my-table"
  mapping_rule_name    = "my-table-mapping"
  data_format          = "JSON"
  retrieval_start_date = "2023-06-26T12:00:00Z"
}
