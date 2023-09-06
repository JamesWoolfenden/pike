data "aws_instances" "pike" {}

output "instances" {
  value = data.aws_instances.pike
}
