resource "aws_vpc_ipam_pool" "example" {
  description    = "this"
  address_family = "ipv4"
  ipam_scope_id  = aws_vpc_ipam.example.private_default_scope_id
  locale         = data.aws_region.current.name
  tags = {
    pike = "permissions"
  }
}
