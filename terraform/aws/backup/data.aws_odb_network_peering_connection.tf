data "aws_odb_network_peering_connection" "pike" {
  id = "pike"
}

output "aws_odb_network_peering_connection" {
  value = data.aws_odb_network_peering_connection.pike
}
