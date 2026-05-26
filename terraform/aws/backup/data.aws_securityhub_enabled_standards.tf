data "aws_securityhub_enabled_standards" "pike" {
}

output "aws_securityhub_enabled_standards" {
  value = data.aws_securityhub_enabled_standards.pike
}
