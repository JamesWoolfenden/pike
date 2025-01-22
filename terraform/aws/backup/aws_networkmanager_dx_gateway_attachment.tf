resource "aws_networkmanager_dx_gateway_attachment" "pike" {
  core_network_id            = aws_networkmanager_core_network_policy_attachment.test.core_network_id
  direct_connect_gateway_arn = "arn:aws:directconnect::${data.aws_caller_identity.current.account_id}:dx-gateway/${aws_dx_gateway.test.id}"
  edge_locations             = [data.aws_region.current.name]
}
