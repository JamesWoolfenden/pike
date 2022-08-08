resource "aws_db_instance" "default" {
  allocated_storage                   = 10
  allow_major_version_upgrade         = true
  auto_minor_version_upgrade          = true
  backup_retention_period             = 7
  engine                              = "mysql"
  engine_version                      = "5.7"
  instance_class                      = "db.t3.micro"
  db_name                             = "mydb"
  username                            = "foo"
  password                            = "foobarbaz"
  parameter_group_name                = "default.mysql5.7"
  skip_final_snapshot                 = true
  availability_zone                   = "eu-west-2a"
  db_subnet_group_name                = "default-vpc-0c33dc8cd64f408c4"
  backup_window                       = "09:46-10:16"
  copy_tags_to_snapshot               = true
  enabled_cloudwatch_logs_exports     = ["audit", "error", "general", "slowquery"]
  iam_database_authentication_enabled = true
  //monitoring_interval = 60
  storage_encrypted = true

  //kms_key_id = ""
  tags = {
    createdby = "JamesWoolfenden"
  }
}
