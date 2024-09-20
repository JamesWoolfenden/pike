resource "aws_network_acl_association" "pike" {
  network_acl_id = aws_network_acl.pike.id
  subnet_id      = aws_subnet.pike.id
}
