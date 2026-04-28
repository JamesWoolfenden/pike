resource "azurerm_netapp_snapshot_policy" "pike_gen" {
  name                = "snapshotpolicy-01"
  location            = "pike"
  resource_group_name = "pike"
  account_name        = "pike"
  enabled             = true

  hourly_schedule {
    snapshots_to_keep = 4
    minute            = 15
  }

  daily_schedule {
    snapshots_to_keep = 2
    hour              = 20
    minute            = 15
  }

  weekly_schedule {
    snapshots_to_keep = 1
    days_of_week      = ["Monday", "Friday"]
    hour              = 23
    minute            = 0
  }

  monthly_schedule {
    snapshots_to_keep = 1
    days_of_month     = [1, 15, 20, 30]
    hour              = 5
    minute            = 45
  }
}
