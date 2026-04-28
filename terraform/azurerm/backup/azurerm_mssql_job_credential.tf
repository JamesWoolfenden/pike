resource "azurerm_mssql_job_credential" "pike_gen" {
  name         = "example-credential"
  job_agent_id = "pike"
  username     = "my-username"
}
