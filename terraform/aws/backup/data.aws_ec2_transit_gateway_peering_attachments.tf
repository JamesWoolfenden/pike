data "aws_ec2_transit_gateway_peering_attachments" "pike" {
}

output "aws_ec2_transit_gateway_peering_attachments" {
  value = data.aws_ec2_transit_gateway_peering_attachments.pike
}
