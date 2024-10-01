data "aws_elasticache_reserved_cache_node_offering" "pike" {
  cache_node_type     = "cache. t4g. small"
  duration            = "P1Y"
  offering_type       = "No Upfront"
  product_description = "redis"
}

output "aws_elasticache_reserved_cache_node_offering" {
  value = data.aws_elasticache_reserved_cache_node_offering.pike
}
