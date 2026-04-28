resource "azurerm_synapse_workspace_security_alert_policy" "pike_gen" {
  synapse_workspace_id = "pike"
  policy_state         = "Enabled"
  storage_endpoint     = "pike"
  disabled_alerts = [
    "Sql_Injection",
    "Data_Exfiltration"
  ]
  retention_days = 20
}
