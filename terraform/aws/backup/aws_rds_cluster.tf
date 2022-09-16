resource "aws_rds_cluster" "pike" {
  allow_major_version_upgrade         = true
  apply_immediately                   = true
  db_cluster_parameter_group_name     = "default.aurora5.6"
  db_instance_parameter_group_name    = "default.aurora5.6"
  db_subnet_group_name                = "default-vpc-0c33dc8cd64f408c4"
  iam_database_authentication_enabled = true
  //iam_roles                           = ["arn:aws:iam::680235478471:role/aws-service-role/rds.amazonaws.com/AWSServiceRoleForRDS"]
  kms_key_id                      = ""
  cluster_identifier              = "aurora-cluster-pike"
  enabled_cloudwatch_logs_exports = ["audit"]
  //db_cluster_instance_class = ""
  engine                    = "aurora"
  engine_version            = "5.6.mysql_aurora.1.22.2"
  availability_zones        = ["eu-west-2a", "eu-west-2b", "eu-west-2c"]
  global_cluster_identifier = "global-test"
  database_name             = "pike"
  master_username           = "foo"
  master_password           = "barlonger"
  backup_retention_period   = 5
  preferred_backup_window   = "07:00-09:00"
  skip_final_snapshot       = true
  //scaling_configuration {}
  storage_encrypted = true
  //vpc_security_group_ids = []
  tags = {
    pike = "permission"
    #    delete="me"
  }
}
