resource "azurerm_automanage_configuration" "pike_gen" {
  name                = "example-acmp"
  resource_group_name = "pike"
  location            = "pike"

  antimalware {
    exclusions {
      extensions = "exe;dll"
      paths      = "C:\\Windows\\Temp;D:\\Temp"
      processes  = "svchost.exe;notepad.exe"
    }

    real_time_protection_enabled   = true
    scheduled_scan_enabled         = true
    scheduled_scan_type            = "Quick"
    scheduled_scan_day             = 1
    scheduled_scan_time_in_minutes = 1339
  }

  azure_security_baseline {
    assignment_type = "ApplyAndAutoCorrect"
  }

  automation_account_enabled = true

  backup {
    policy_name                        = "acctest-backup-policy-%d"
    time_zone                          = "UTC"
    instant_rp_retention_range_in_days = 2

    schedule_policy {
      schedule_run_frequency = "Daily"
      schedule_run_days      = ["Monday", "Tuesday"]
      schedule_run_times     = ["12:00"]
      schedule_policy_type   = "SimpleSchedulePolicy"
    }

    retention_policy {
      retention_policy_type = "LongTermRetentionPolicy"

      daily_schedule {
        retention_times = ["12:00"]
        retention_duration {
          count         = 7
          duration_type = "Days"
        }
      }

      weekly_schedule {
        retention_times = ["14:00"]
        retention_duration {
          count         = 4
          duration_type = "Weeks"
        }
      }
    }
  }

  boot_diagnostics_enabled    = true
  defender_for_cloud_enabled  = true
  guest_configuration_enabled = true
  log_analytics_enabled       = true
  status_change_alert_enabled = true

  tags = {
    "env" = "test"
  }
}
