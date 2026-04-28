resource "azurerm_datadog_monitor_sso_configuration" "pike_gen" {
  datadog_monitor_id        = "pike"
  single_sign_on            = "Enable"
  enterprise_application_id = "00000000-0000-0000-0000-000000000000"
}
