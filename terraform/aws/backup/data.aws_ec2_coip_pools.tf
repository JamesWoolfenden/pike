data "aws_ec2_coip_pools" "pike" {}

output "pools" {
  value = data.aws_ec2_coip_pools.pike
}
