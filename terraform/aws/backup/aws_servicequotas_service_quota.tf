resource "aws_servicequotas_service_quota" "pike" {
  quota_code   = "L-F678F1CE"
  service_code = "vpc"
  value        = 5
}
