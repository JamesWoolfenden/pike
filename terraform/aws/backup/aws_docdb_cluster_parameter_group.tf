resource "aws_docdb_cluster_parameter_group" "pike" {
  family      = "docdb3.6"
  name        = "pike"
  description = "docdb cluster parameter group"

  parameter {
    name  = "tls"
    value = "enabled"
  }

  tags = {
    pike   = "permissions"
    delete = "me"
  }
}
