data "aws_datazone_domain" "pike" {
}

output "aws_datazone_domain" {
  value = data.aws_datazone_domain.pike
}
