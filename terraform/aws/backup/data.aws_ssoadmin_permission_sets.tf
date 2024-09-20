data "aws_ssoadmin_instances" "example" {}

data "aws_ssoadmin_permission_sets" "pike" {
  instance_arn = tolist(data.aws_ssoadmin_instances.example.arns)[0]
}

output "aws_ssoadmin_permission_sets" {
  value = data.aws_ssoadmin_permission_sets.pike
}
