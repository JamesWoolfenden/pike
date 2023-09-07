data "aws_ec2_public_ipv4_pools" "pike" {}

output "pools" {
  value = data.aws_ec2_public_ipv4_pools.pike
}
