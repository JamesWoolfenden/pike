data "aws_ec2_instance_types" "pike" {}

output "types" {
  value = data.aws_ec2_instance_types.pike
}
