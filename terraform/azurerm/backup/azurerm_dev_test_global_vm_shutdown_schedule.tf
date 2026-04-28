resource "azurerm_dev_test_global_vm_shutdown_schedule" "pike_gen" {
  virtual_machine_id = "pike"
  location           = "pike"
  enabled            = true

  daily_recurrence_time = "1100"
  timezone              = "Pacific Standard Time"

  notification_settings {
    enabled         = true
    time_in_minutes = "60"
    webhook_url     = "https://sample-webhook-url.example.com"
  }
}
