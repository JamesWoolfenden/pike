resource "azurerm_mssql_job_agent" "pike_gen" {
  name        = "example-job-agent"
  location    = "pike"
  database_id = "pike"
}
