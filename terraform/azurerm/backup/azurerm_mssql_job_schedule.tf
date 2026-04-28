resource "azurerm_mssql_job_schedule" "pike_gen" {
  job_id = "pike"

  type       = "Recurring"
  enabled    = true
  end_time   = "2025-12-01T00:00:00Z"
  interval   = "PT5M"
  start_time = "2025-01-01T00:00:00Z"
}
