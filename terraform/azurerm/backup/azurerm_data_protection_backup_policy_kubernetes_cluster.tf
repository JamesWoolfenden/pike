resource "azurerm_data_protection_backup_policy_kubernetes_cluster" "pike_gen" {
  name                = "example-backup-policy"
  resource_group_name = "pike"
  vault_name          = "pike"

  backup_repeating_time_intervals = ["R/2021-05-23T02:30:00+00:00/P1W"]
  time_zone                       = "India Standard Time"
  default_retention_duration      = "P4M"

  retention_rule {
    name     = "Daily"
    priority = 25

    life_cycle {
      duration        = "P84D"
      data_store_type = "OperationalStore"
    }

    criteria {
      absolute_criteria = "FirstOfDay"
    }
  }

  default_retention_rule {
    life_cycle {
      duration        = "P7D"
      data_store_type = "OperationalStore"
    }
  }
}
