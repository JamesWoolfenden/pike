data "aws_vpc_peering_connections" "pike" {}

output "conns" {
  value = data.aws_vpc_peering_connections.pike
}
