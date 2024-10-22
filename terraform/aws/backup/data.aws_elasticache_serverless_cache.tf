data "aws_elasticache_serverless_cache" "pike" {
  name = "pike"
}

output "aws_elasticache_serverless_cache" {
  value = data.aws_elasticache_serverless_cache.pike
}
