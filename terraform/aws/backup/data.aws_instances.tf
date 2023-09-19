data "aws_instances" "pike" {}

output "aws_instances" {
  value = data.aws_instances.pike
}
