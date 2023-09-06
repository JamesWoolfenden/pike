data "aws_subnets" "pike" {}

output "subnets" {
  value = data.aws_subnets.pike
}
