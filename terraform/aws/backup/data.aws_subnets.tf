data "aws_subnets" "pike" {}

output "aws_subnets" {
  value = data.aws_subnets.pike
}
