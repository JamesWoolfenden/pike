resource "aws_elasticache_parameter_group" "pike" {
  name   = "cache-params"
  family = "redis2.8"

  parameter {
    name  = "activerehashing"
    value = "yes"
  }

  parameter {
    name  = "min-slaves-to-write"
    value = "2"
  }
  tags = {
    pike = "permissions"
    //delete="me"
  }
}
