data "aws_nat_gateways" "pike" {}

output "nat" {
  value = data.aws_nat_gateways.pike
}
