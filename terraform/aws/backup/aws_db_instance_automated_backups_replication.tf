resource "aws_db_instance_automated_backups_replication" "pike" {
  provider               = aws.central
  source_db_instance_arn = aws_db_instance.pike.arn
  retention_period       = 14
}
