data "aws_ssoadmin_instances" "example" {}

data "aws_ssoadmin_permission_set" "example" {
  instance_arn = tolist(data.aws_ssoadmin_instances.example.arns)[0]
  name         = "Example"
}

output "arn" {
  value = data.aws_ssoadmin_permission_set.example.arn
}
