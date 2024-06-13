resource "aws_ec2_host" "pike" {
  availability_zone = "eu-west-2a"
  host_recovery     = "on"
  auto_placement    = "on"
  instance_type     = "t3.micro"
  tags = {
    pike = "permissions"
  }
}
