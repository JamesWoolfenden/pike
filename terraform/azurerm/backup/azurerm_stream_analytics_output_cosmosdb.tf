resource "azurerm_stream_analytics_output_cosmosdb" "pike_gen" {
  name                     = "output-to-cosmosdb"
  stream_analytics_job_id  = "pike"
  cosmosdb_account_key     = "pike"
  cosmosdb_sql_database_id = "pike"
  container_name           = "pike"
  document_id              = "exampledocumentid"
}
