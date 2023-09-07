data "aws_ssm_maintenance_windows" "pike" {}

output "windows" {
  value = data.aws_ssm_maintenance_windows.pike
}