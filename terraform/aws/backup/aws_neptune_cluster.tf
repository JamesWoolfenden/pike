resource "aws_neptune_cluster" "pike" {
  neptune_subnet_group_name            = "main"
  allow_major_version_upgrade          = true
  engine_version                       = "1.2.0.0"
  neptune_cluster_parameter_group_name = "dev-appname-neptune-cluster-parameter-group"
  cluster_identifier_prefix            = "neptune-cluster-demo"
  engine                               = "neptune"
  backup_retention_period              = 5
  preferred_backup_window              = "07:00-09:00"
  skip_final_snapshot                  = true
  iam_database_authentication_enabled  = true
  apply_immediately                    = true
  copy_tags_to_snapshot                = true
  enable_cloudwatch_logs_exports       = ["audit"]
  final_snapshot_identifier            = "sato"
  iam_roles                            = ["arn:aws:iam::680235478471:role/dev-appname-neptune-iam-role-eu-west-2"]
  vpc_security_group_ids               = ["sg-018be3df336b6ae00"]
}
