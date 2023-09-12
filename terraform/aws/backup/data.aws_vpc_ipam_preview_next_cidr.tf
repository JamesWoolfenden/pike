data "aws_vpc_ipam_preview_next_cidr" "test" {
  ipam_pool_id   = aws_vpc_ipam_pool.example.id
  netmask_length = 28

}

#resource "aws_vpc_ipam_pool_cidr_allocation" "test" {
#  ipam_pool_id = aws_vpc_ipam_pool.example.id
#  cidr         = data.aws_vpc_ipam_preview_next_cidr.test.cidr
#
#  lifecycle {
#    ignore_changes = [cidr]
#  }
#}
#
#data "aws_region" "current" {}
#
#resource "aws_vpc_ipam" "example" {
#  operating_regions {
#    region_name = data.aws_region.current.name
#  }
#}
#
#resource "aws_vpc_ipam_pool" "example" {
#  address_family = "ipv4"
#  ipam_scope_id  = aws_vpc_ipam.example.private_default_scope_id
#  locale         = data.aws_region.current.name
#}
