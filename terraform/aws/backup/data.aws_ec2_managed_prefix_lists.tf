data "aws_ec2_managed_prefix_lists" "pike" {}

output "list" {
  value = data.aws_ec2_managed_prefix_lists.pike
}
