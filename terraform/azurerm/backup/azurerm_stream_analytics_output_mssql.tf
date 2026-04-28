resource "azurerm_stream_analytics_output_mssql" "pike_gen" {
  name                      = "example-output-sql"
  stream_analytics_job_name = "pike"
  resource_group_name       = "pike"

  server   = "pike"
  user     = "pike"
  database = "pike"
  table    = "ExampleTable"
}
