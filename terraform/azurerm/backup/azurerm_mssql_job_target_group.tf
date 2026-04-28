resource "azurerm_mssql_job_target_group" "pike_gen" {
  name         = "example-target-group"
  job_agent_id = "pike"

  job_target {
    server_name       = "pike"
    job_credential_id = "pike"
  }
}
