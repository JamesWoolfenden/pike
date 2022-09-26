resource "aws_vpc_ipv4_cidr_block_association" "pike" {
  vpc_id     = "vpc-0c33dc8cd64f408c4"
  cidr_block = "172.2.0.0/16"
}
