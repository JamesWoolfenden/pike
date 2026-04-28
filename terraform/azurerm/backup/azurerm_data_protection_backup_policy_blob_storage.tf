resource "azurerm_data_protection_backup_policy_blob_storage" "pike_gen" {
  name                                   = "example-backup-policy"
  vault_id                               = "pike"
  operational_default_retention_duration = "P30D"
  vault_default_retention_duration       = "P7D"

  retention_rule {
    name     = "Weekly"
    priority = 20

    life_cycle {
      duration        = "P90D"
      data_store_type = "VaultStore"
    }

    criteria {
      days_of_week = ["Monday"]
    }
  }

  retention_rule {
    name     = "Monthly"
    priority = 10

    life_cycle {
      duration        = "P180D"
      data_store_type = "VaultStore"
    }

    criteria {
      days_of_month = [1]
    }
  }

  retention_rule {
    name     = "Yearly"
    priority = 5

    life_cycle {
      duration        = "P365D"
      data_store_type = "VaultStore"
    }

    criteria {
      months_of_year = ["January"]
      days_of_month  = [1]
    }
  }
}
