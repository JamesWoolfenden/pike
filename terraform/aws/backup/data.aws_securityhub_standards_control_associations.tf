data "aws_securityhub_standards_control_associations" "pike" {
  security_control_id = "pike"
}

output "aws_securityhub_standards_control_associations" {
  value = data.aws_securityhub_standards_control_associations.pike
}
