resource "aws_vpc_ipam_preview_next_cidr" "example" {
  ipam_pool_id   = aws_vpc_ipam_pool.example.id
  netmask_length = 28

  disallowed_cidrs = [
    "172.2.0.0/32",
  ]

  depends_on = [
    aws_vpc_ipam_pool_cidr.example
  ]

}
