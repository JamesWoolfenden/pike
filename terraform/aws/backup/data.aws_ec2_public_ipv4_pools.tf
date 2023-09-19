data "aws_ec2_public_ipv4_pools" "pike" {}

output "ipv4_pools" {
  value = data.aws_ec2_public_ipv4_pools.pike
}
