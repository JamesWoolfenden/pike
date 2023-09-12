data "aws_ssm_instances" "pike" {}

output "instances" {
  value = data.aws_ssm_instances.pike
}
