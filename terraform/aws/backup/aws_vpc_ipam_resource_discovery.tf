resource "aws_vpc_ipam_resource_discovery" "pike" {
  description = "My IPAM Resource Discovery"
  operating_regions {
    region_name = data.aws_region.current.name
  }

  tags = {
    another = "tag"
  }
}
