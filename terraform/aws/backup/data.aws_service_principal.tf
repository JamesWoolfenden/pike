data "aws_service_principal" "pike" {
  service_name = "s3"
}

output "aws_service_principal" {
  value = data.aws_service_principal.pike
}
