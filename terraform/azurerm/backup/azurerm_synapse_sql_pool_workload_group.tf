resource "azurerm_synapse_sql_pool_workload_group" "pike_gen" {
  name                               = "example"
  sql_pool_id                        = "pike"
  importance                         = "normal"
  max_resource_percent               = 100
  min_resource_percent               = 0
  max_resource_percent_per_request   = 3
  min_resource_percent_per_request   = 3
  query_execution_timeout_in_seconds = 0
}
