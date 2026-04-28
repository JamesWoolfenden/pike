resource "azurerm_mssql_job" "pike_gen" {
  name         = "example-job"
  job_agent_id = "pike"
  description  = "example description"
}
