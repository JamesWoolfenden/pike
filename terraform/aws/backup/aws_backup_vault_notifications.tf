resource "aws_backup_vault_notifications" "pike" {
  backup_vault_name   = aws_backup_vault.pike.name
  sns_topic_arn       = "arn:aws:sns:eu-west-2:680235478471:pike"
  backup_vault_events = ["BACKUP_JOB_STARTED", "RESTORE_JOB_COMPLETED"]
}
