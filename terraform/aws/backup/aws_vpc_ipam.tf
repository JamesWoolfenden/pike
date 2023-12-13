resource "aws_vpc_ipam" "example" {
  description = "this"
  operating_regions {
    region_name = data.aws_region.current.name
  }
  tags = {
    pike    = "permissions"
    another = "tag"
  }
}
