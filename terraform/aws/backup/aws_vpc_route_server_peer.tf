# resource "aws_vpc_route_server_peer" "pike" {
#   provider                 = aws.central
#   route_server_endpoint_id = aws_vpc_route_server_endpoint.pike.route_server_endpoint_id
#   peer_address             = "10.0.1.250"
#   bgp_options {
#     peer_asn = 65200
#   }
#
#   tags = {
#     Name = "Appliance 1"
#   }
# }
