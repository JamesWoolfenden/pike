data "aws_ec2_instance_type" "pike" {
  instance_type = "t3.micro"
}

output "type" {
  value = data.aws_ec2_instance_type.pike
}
