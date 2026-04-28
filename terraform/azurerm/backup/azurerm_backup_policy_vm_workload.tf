resource "azurerm_backup_policy_vm_workload" "pike_gen" {
  name                = "example-bpvmw"
  resource_group_name = "pike"
  recovery_vault_name = "pike"

  workload_type = "SQLDataBase"

  settings {
    time_zone           = "UTC"
    compression_enabled = false
  }

  protection_policy {
    policy_type = "Full"

    backup {
      frequency = "Daily"
      time      = "15:00"
    }

    retention_daily {
      count = 8
    }
  }

  protection_policy {
    policy_type = "Log"

    backup {
      frequency_in_minutes = 15
    }

    simple_retention {
      count = 8
    }
  }
}
