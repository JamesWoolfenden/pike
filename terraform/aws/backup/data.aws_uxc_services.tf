data "aws_uxc_services" "pike" {
}

output "aws_uxc_services" {
  value = data.aws_uxc_services.pike
}
