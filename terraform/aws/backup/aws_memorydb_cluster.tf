resource "aws_memorydb_cluster" "pike" {
  acl_name                 = "open-access"
  name                     = "pike"
  node_type                = "db.t4g.small"
  num_shards               = 2
  security_group_ids       = ["sg-05b27cb61c9c46bd2"]
  snapshot_retention_limit = 7
  subnet_group_name        = "pike"
  tags = {
    pike = "permissions"
  }
}
