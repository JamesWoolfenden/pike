resource "aws_nat_gateway" "pike" {
  subnet_id         = "subnet-0562ef1d304b968f4"
  allocation_id     = "eipalloc-0047fa56c40637c3b"
  connectivity_type = "public"
  tags = {
    pike = "permissions"
  }
}
