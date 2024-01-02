resource "aws_verifiedaccess_group" "pike" {
  verifiedaccess_instance_id = aws_verifiedaccess_instance.pike.id

  tags = {
    pike = "permissions"
  }
}
