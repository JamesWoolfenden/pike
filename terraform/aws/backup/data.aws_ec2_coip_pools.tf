data "aws_ec2_coip_pools" "pike" {}

output "coip_pools" {
  value = data.aws_ec2_coip_pools.pike
}
