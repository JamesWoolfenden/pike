data "aws_odb_network_peering_connections" "pike" {

}

output "aws_odb_network_peering_connections" {
  value = data.aws_odb_network_peering_connections.pike
}
