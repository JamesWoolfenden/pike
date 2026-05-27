data "aws_msk_topic" "pike" {
  cluster_arn = "arn:aws:kafka:eu-west-2:123456789012:cluster/pike/00000000-0000-0000-0000-000000000000-0"
  name        = "pike"
}

output "aws_msk_topic" {
  value = data.aws_msk_topic.pike
}
