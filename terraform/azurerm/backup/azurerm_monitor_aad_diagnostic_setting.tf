resource "azurerm_monitor_aad_diagnostic_setting" "pike_gen" {
  name               = "setting1"
  storage_account_id = "pike"
  enabled_log {
    category = "SignInLogs"
  }
  enabled_log {
    category = "AuditLogs"
  }
  enabled_log {
    category = "NonInteractiveUserSignInLogs"
  }
  enabled_log {
    category = "ServicePrincipalSignInLogs"
  }
}
