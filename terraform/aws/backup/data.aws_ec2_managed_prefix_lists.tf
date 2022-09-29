data "aws_region" "current" {}

data "aws_ec2_managed_prefix_list" "example" {
  name = "com.amazonaws.${data.aws_region.current.name}.dynamodb"
}

output "prefix_list" {
  value = data.aws_ec2_managed_prefix_list.example
}
