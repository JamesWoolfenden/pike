data "aws_servicequotas_service" "pike" {
  service_name = "Amazon Virtual Private Cloud (Amazon VPC)"
}

output "quota" {
  value=data.aws_servicequotas_service.pike
}
//servicequotas:ListServices