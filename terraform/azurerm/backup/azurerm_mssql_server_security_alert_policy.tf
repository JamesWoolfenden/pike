resource "azurerm_mssql_server_security_alert_policy" "pike_gen" {
  resource_group_name = "pike"
  server_name         = "pike"
  state               = "Enabled"
  storage_endpoint    = "pike"
  retention_days      = 20

  disabled_alerts = [
    "Sql_Injection",
    "Data_Exfiltration"
  ]
}
