data "aws_securityhub_security_controls" "pike" {
}

output "aws_securityhub_security_controls" {
  value = data.aws_securityhub_security_controls.pike
}
