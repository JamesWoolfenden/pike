resource "aws_ami_from_instance" "pike" {
  name               = "pike-from-instance"
  source_instance_id = "i-0eac08fa6c940b0d0"
  tags = {
    pike = "permissions"
  }
}

output "ami2" {
  value = aws_ami_from_instance.pike
}
