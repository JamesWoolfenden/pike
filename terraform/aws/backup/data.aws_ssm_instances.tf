data "aws_ssm_instances" "pike" {}

output "ssm_instances" {
  value = data.aws_ssm_instances.pike
}
