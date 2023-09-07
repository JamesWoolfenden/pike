data "aws_ec2_instance_type_offerings" "pike" {}

output "offerings" {
  value = data.aws_ec2_instance_type_offerings.pike
}
