resource "aws_ami_copy" "pike" {
  name              = "pike2"
  source_ami_id     = aws_ami.pike.id
  source_ami_region = "eu-west-2"
  tags = {
    pike = "permissions"
  }
}
