resource "azurerm_monitor_alert_processing_rule_action_group" "pike_gen" {
  name                 = "example"
  resource_group_name  = "example"
  scopes               = ["pike"]
  add_action_group_ids = ["pike"]

  condition {
    target_resource_type {
      operator = "Equals"
      values   = ["Microsoft.Compute/VirtualMachines"]
    }
    severity {
      operator = "Equals"
      values   = ["Sev0", "Sev1", "Sev2"]
    }
  }

  schedule {
    effective_from  = "2022-01-01T01:02:03"
    effective_until = "2022-02-02T01:02:03"
    time_zone       = "Pacific Standard Time"
    recurrence {
      daily {
        start_time = "17:00:00"
        end_time   = "09:00:00"
      }
      weekly {
        days_of_week = ["Saturday", "Sunday"]
      }
    }
  }

  tags = {
    foo = "bar"
  }
}
