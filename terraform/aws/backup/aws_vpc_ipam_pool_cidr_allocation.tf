data "aws_region" "current" {}
#
resource "aws_vpc_ipam_pool_cidr_allocation" "example" {
  ipam_pool_id = aws_vpc_ipam_pool.example.id
  cidr         = "172.20.0.0/24"
  depends_on = [
    aws_vpc_ipam_pool_cidr.example
  ]
  description = "pike"
}
