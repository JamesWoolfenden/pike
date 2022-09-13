resource "aws_redshift_parameter_group" "pike" {
  name   = "parameter-group-pike"
  family = "redshift-1.0"

  parameter {
    name  = "require_ssl"
    value = "false"
  }

  parameter {
    name  = "query_group"
    value = "example"
  }

  parameter {
    name  = "enable_user_activity_logging"
    value = "true"
  }

  tags = {
    pike = "permission"
    //delete="me"
  }
}
