resource "aws_backup_global_settings" "pike" {
  global_settings = {
    "isCrossAccountBackupEnabled" = "true"
  }
}
