resource "azurerm_oracle_autonomous_database_backup" "pike_gen" {
  name                     = "example-backup"
  autonomous_database_id   = "pike"
  retention_period_in_days = 120
  backup_type              = "Full"
}
