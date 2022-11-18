resource "aws_backup_vault_lock_configuration" "pike" {
  backup_vault_name   = aws_backup_vault.pike.name
  changeable_for_days = 3
  max_retention_days  = 1200
  min_retention_days  = 7
}
