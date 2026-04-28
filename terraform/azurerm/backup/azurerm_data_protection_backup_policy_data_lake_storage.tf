resource "azurerm_data_protection_backup_policy_data_lake_storage" "pike_gen" {
  name                            = "example-backup-policy"
  data_protection_backup_vault_id = "pike"
  backup_schedule                 = ["R/2021-05-23T02:30:00+00:00/P1W"]
  time_zone                       = "India Standard Time"

  default_retention_duration = "P4M"

  retention_rule {
    name              = "weekly"
    duration          = "P6M"
    absolute_criteria = "FirstOfWeek"
  }

  retention_rule {
    name                   = "thursday"
    duration               = "P1W"
    days_of_week           = ["Thursday"]
    scheduled_backup_times = ["2021-05-23T02:30:00Z"]
  }

  retention_rule {
    name                   = "monthly"
    duration               = "P1D"
    weeks_of_month         = ["First", "Last"]
    days_of_week           = ["Tuesday"]
    scheduled_backup_times = ["2021-05-23T02:30:00Z"]
  }
}
