data "aws_ec2_transit_gateway_route_tables" "pike" {}

output "tables" {
  value = data.aws_ec2_transit_gateway_route_tables.pike
}
