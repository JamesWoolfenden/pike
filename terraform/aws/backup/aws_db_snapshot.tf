resource "aws_db_instance" "pike" {
  allocated_storage = 10
  engine            = "mysql"
  engine_version    = "8.0.37"
  instance_class    = "db.t3.micro"
  db_name           = "baz"
  password          = "barbarbarbar"
  username          = "foo"

  maintenance_window      = "Fri:09:00-Fri:09:30"
  backup_retention_period = 0
  parameter_group_name    = "default.mysql8.0"
  skip_final_snapshot     = true
}

resource "aws_db_snapshot" "pike" {
  db_instance_identifier = aws_db_instance.pike.identifier
  db_snapshot_identifier = "testsnapshot1234"
}
