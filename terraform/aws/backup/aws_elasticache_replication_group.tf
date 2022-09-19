resource "aws_elasticache_replication_group" "pike" {
  automatic_failover_enabled  = true
  preferred_cache_cluster_azs = ["eu-west-2a", "eu-west-2b"]
  replication_group_id        = "tf-rep-group-1"
  description                 = "example description"
  node_type                   = "cache.m4.large"
  num_cache_clusters          = 2
  parameter_group_name        = "default.redis6.x"
  port                        = 6379

  tags = {
    pike   = "permission"
    delete = "me"
  }
}
