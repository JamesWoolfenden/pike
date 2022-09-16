resource "aws_efs_backup_policy" "pike" {
  backup_policy {
    status = "DISABLED"
  }
  file_system_id = "fs-0898bc857b16b617a"
}
