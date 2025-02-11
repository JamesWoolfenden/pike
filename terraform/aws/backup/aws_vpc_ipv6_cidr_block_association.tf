resource "aws_vpc_ipv6_cidr_block_association" "pike" {
  ipv6_ipam_pool_id   = aws_vpc_ipam_pool.example.id
  vpc_id              = aws_vpc.example.id
  ipv6_netmask_length = 60
}

data "aws_region" "current" {}

resource "aws_vpc_ipam" "example" {
  operating_regions {
    region_name = data.aws_region.current.name
  }
}

resource "aws_vpc_ipam_pool" "example" {
  address_family = "ipv4"
  ipam_scope_id  = aws_vpc_ipam.example.private_default_scope_id
  locale         = data.aws_region.current.name
}
