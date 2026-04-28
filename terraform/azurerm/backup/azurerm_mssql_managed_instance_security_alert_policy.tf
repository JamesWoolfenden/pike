resource "azurerm_mssql_managed_instance_security_alert_policy" "pike_gen" {
  resource_group_name   = "pike"
  managed_instance_name = "pike"
  enabled               = true
  storage_endpoint      = "pike"
  disabled_alerts = [
    "Sql_Injection",
    "Data_Exfiltration"
  ]
  retention_days = 20
}
