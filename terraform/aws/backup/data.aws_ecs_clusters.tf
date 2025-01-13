data "aws_ecs_clusters" "pike" {
}

output "aws_ecs_clusters" {
  value = data.aws_ecs_clusters.pike
}
