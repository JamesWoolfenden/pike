resource "google_netapp_backup_policy" "pike" {
  name                 = "test-backup-policy-full"
  location             = "us-central1"
  daily_backup_limit   = 2
  weekly_backup_limit  = 1
  monthly_backup_limit = 1
  description          = "TF test backup schedule"
  enabled              = true
  labels = {
    "foo" = "bar"
  }
}
