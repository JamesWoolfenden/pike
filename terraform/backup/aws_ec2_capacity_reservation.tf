resource "aws_ec2_capacity_reservation" "default" {
  instance_type     = "t3a.micro"
  instance_platform = "Linux/UNIX"
  availability_zone = "eu-west-2a"
  ephemeral_storage = false
  ebs_optimized     = true
  instance_count    = 1
  tags = {
    pike = "permissions"
  }
}
