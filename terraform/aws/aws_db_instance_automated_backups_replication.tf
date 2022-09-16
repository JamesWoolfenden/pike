#resource "aws_db_instance_automated_backups_replication" "pike" {
#  source_db_instance_arn = aws_rds_cluster_instance.pike.arn
#  //todo:probably needs to be a checkov test
#  retention_period = 0
#  //kms_key_id = ""
#}
