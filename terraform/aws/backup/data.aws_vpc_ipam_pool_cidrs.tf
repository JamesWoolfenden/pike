data "aws_vpc_ipam_pool_cidrs" "pike" {
  ipam_pool_id = data.aws_vpc_ipam_pool.p.id
}

data "aws_vpc_ipam_pool" "p" {
  filter {
    name   = "description"
    values = ["*mypool*"]
  }

  filter {
    name   = "address-family"
    values = ["ipv4"]
  }
}
