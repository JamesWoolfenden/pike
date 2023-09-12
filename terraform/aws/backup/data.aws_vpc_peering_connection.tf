data "aws_vpc_peering_connection" "pike" {}

output "conn" {
  value = data.aws_vpc_peering_connection.pike
}
