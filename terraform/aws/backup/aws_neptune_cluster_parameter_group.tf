resource "aws_neptune_cluster_parameter_group" "pike" {
  family      = "neptune1"
  name        = "example"
  description = "neptune cluster parameter group"

  parameter {
    name  = "neptune_enable_audit_log"
    value = 1
  }
  tags = {
    pike = "permissions"
    // delete="me"
  }
}
