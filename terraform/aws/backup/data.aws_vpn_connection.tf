data "aws_vpn_connection" "pike" {
  vpn_connection_id = "pike"
}

output "aws_vpn_connection" {
  value = data.aws_vpn_connection.pike
}
