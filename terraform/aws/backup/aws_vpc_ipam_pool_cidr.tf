resource "aws_vpc_ipam_pool_cidr" "example" {
  ipam_pool_id = aws_vpc_ipam_pool.example.id
  cidr         = "172.20.0.0/16"
}
