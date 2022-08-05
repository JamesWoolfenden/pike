resource "aws_mq_broker" "broker" {
  broker_name                = "mybrokername"
  auto_minor_version_upgrade = true
  apply_immediately          = true
  authentication_strategy    = "ldap"

  configuration {
    id       = aws_mq_configuration.broker.id
    revision = aws_mq_configuration.broker.latest_revision
  }

  engine_type         = "ActiveMQ" //RABBITMQ
  engine_version      = "5.15.9"
  host_instance_type  = "mq.t2.micro"
  deployment_mode     = "SINGLE_INSTANCE"
  publicly_accessible = true
  security_groups     = ["sg-05b27cb61c9c46bd2"]
  storage_type        = "efs"
  #
  user {
    username = "Fred"
    password = "QuimbyWasAGod"
  }

  user {
    username = "Sally"
    password = "QuimbyWasAGod"
  }

  maintenance_window_start_time {
    day_of_week = "MONDAY"
    time_of_day = "12:05"
    time_zone   = "GMT"
  }

  encryption_options {
    // kms_key_id        = var.kms_key_id
    use_aws_owned_key = true
  }

  logs {
    general = false
    audit   = false
  }

  subnet_ids = ["subnet-08d97e381dbc80d40"]
  tags       = { name = "some_tags" }
}
