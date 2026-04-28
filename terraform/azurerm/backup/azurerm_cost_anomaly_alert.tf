resource "azurerm_cost_anomaly_alert" "pike_gen" {
  name            = "alertname"
  display_name    = "Alert DisplayName"
  subscription_id = "/subscriptions/00000000-0000-0000-0000-000000000000"
  email_subject   = "My Test Anomaly Alert"
  email_addresses = ["example@test.net"]
}
