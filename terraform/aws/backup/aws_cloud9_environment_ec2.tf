resource "aws_cloud9_environment_ec2" "pike" {
  instance_type               = "t2.micro"
  name                        = "pike-env"
  automatic_stop_time_minutes = 90
  connection_type             = "CONNECT_SSH"
  description                 = "The description of the environment."
  image_id                    = "amazonlinux-1-x86_64"
  owner_arn                   = "arn:aws:iam::680235478471:user/basic"
  subnet_id                   = "subnet-0562ef1d304b968f4"
  tags = {
    pike = "permission"
  }
}

output "cloud9" {
  value = aws_cloud9_environment_ec2.pike
}
