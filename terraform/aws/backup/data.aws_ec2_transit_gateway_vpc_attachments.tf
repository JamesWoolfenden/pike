data "aws_ec2_transit_gateway_vpc_attachments" "pike" {}

output "attachments" {
  value = data.aws_ec2_transit_gateway_vpc_attachments.pike
}
