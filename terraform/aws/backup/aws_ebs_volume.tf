resource "aws_ebs_volume" "example" {
  availability_zone = "eu-west-2a"
  size              = 50
  encrypted         = true
  type              = "gp3"
  #  tags = {
  #    Name = "HelloWorld"
  #    pike="permissions"
  #  }
}
