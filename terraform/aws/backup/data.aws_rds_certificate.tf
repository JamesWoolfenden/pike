data "aws_rds_certificate" "pike" {}

output "cert" {
  value = data.aws_rds_certificate.pike
}
