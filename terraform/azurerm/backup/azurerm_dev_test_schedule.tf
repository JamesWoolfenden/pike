resource "azurerm_dev_test_schedule" "pike_gen" {
  name                = "LabVmAutoStart"
  location            = "pike"
  resource_group_name = "pike"
  lab_name            = "pike"
  status              = "Enabled"

  weekly_recurrence {
    time      = "1100"
    week_days = ["Monday", "Tuesday"]
  }

  time_zone_id = "Pacific Standard Time"
  task_type    = "LabVmsStartupTask"

  notification_settings {
  }

  tags = {
    environment = "Production"
  }
}
