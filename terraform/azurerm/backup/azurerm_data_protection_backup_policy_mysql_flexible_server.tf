resource "azurerm_data_protection_backup_policy_mysql_flexible_server" "pike_gen" {
  name                            = "example-backup-policy"
  vault_id                        = "pike"
  backup_repeating_time_intervals = ["R/2021-05-23T02:30:00+00:00/P1W"]
  time_zone                       = "India Standard Time"

  default_retention_rule {
    life_cycle {
      duration        = "P4M"
      data_store_type = "VaultStore"
    }
  }

  retention_rule {
    name = "weekly"
    life_cycle {
      duration        = "P6M"
      data_store_type = "VaultStore"
    }
    priority = 20

    criteria {
      absolute_criteria = "FirstOfWeek"
    }
  }

  retention_rule {
    name = "thursday"
    life_cycle {
      duration        = "P1W"
      data_store_type = "VaultStore"
    }
    priority = 25

    criteria {
      days_of_week           = ["Thursday"]
      scheduled_backup_times = ["2021-05-23T02:30:00Z"]
    }
  }

  retention_rule {
    name = "monthly"
    life_cycle {
      duration        = "P1D"
      data_store_type = "VaultStore"
    }
    priority = 15

    criteria {
      weeks_of_month         = ["First", "Last"]
      days_of_week           = ["Tuesday"]
      scheduled_backup_times = ["2021-05-23T02:30:00Z"]
    }
  }
}
