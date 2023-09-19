data "aws_ec2_transit_gateway_route_table_routes" "pike" {
  filter {
    name   = "type"
    values = ["propagated"]
  }

  transit_gateway_route_table_id = data.aws_ec2_transit_gateway_route_table.pike.id
}

output "table_routes" {
  value = data.aws_ec2_transit_gateway_route_table_routes.pike
}
