resource "aws_vpc_ipam_resource_discovery_association" "test" {
  ipam_id                    = aws_vpc_ipam.example.id
  ipam_resource_discovery_id = aws_vpc_ipam_resource_discovery.pike.id

  tags = {
    "Name"  = "test"
    pike    = "permissions"
    another = "tag"
  }
}
