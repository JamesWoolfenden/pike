resource "aws_neptune_parameter_group" "pike" {
  family = "neptune1"
  name   = "example"

  parameter {
    name  = "neptune_query_timeout"
    value = "25"
  }

  tags = {
    pike   = "permissions"
    delete = "me"
  }
}
