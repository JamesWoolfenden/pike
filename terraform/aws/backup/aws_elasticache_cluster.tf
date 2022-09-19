resource "aws_elasticache_cluster" "pike" {
  cluster_id      = "pike"
  engine          = "redis"
  node_type       = "cache.t2.micro"
  num_cache_nodes = 1
  //parameter_group_name = "default.redis3.2"
  engine_version = "5.0.6"
  port           = 6379

  tags = {
    pike   = "permission"
    delete = "me"
  }
}
