resource "aws_vpc_route_server" "pike" {
  provider                  = aws.central
  amazon_side_asn           = 65534
  persist_routes            = "enable"
  persist_routes_duration   = 2
  sns_notifications_enabled = true

  tags = {
    Name = "Main Route Server"
  }
}
