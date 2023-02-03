data "aws_servicequotas_service_quota" "by_quota_code" {
  quota_code   = "L-F678F1CE"
  service_code = "vpc"
}

output "service_quota" {
  value = data.aws_servicequotas_service_quota.by_quota_code
}
