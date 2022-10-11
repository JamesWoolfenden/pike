resource "aws_dax_cluster" "pike" {
  cluster_name                     = "pike"
  cluster_endpoint_encryption_type = "TLS"
  iam_role_arn                     = "arn:aws:iam::680235478471:role/DAXServiceRoleForDynamoDBAccess20221011071854910300000001"
  node_type                        = "dax.t2.small"
  replication_factor               = 1
  notification_topic_arn           = "arn:aws:sns:eu-west-2:680235478471:pike"
  parameter_group_name             = ""
  maintenance_window               = "sun:05:00-sun:09:00"
  server_side_encryption {
    enabled = true
  }
  tags = {
    pike = "permissions"
  }
}
