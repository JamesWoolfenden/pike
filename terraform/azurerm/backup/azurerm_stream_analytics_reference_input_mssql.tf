resource "azurerm_stream_analytics_reference_input_mssql" "pike_gen" {
  name                      = "example-reference-input"
  resource_group_name       = "pike"
  stream_analytics_job_name = "pike"
  server                    = "pike"
  database                  = "pike"
  username                  = "exampleuser"
  refresh_type              = "RefreshPeriodicallyWithFull"
  refresh_interval_duration = "00:20:00"
  full_snapshot_query       = <<QUERY
    SELECT *
    INTO [YourOutputAlias]
    FROM [YourInputAlias]
QUERY
}
