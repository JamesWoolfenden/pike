data "aws_caller_identity" "current" {}

resource "aws_vpc_endpoint_service_allowed_principal" "pike" {
  vpc_endpoint_service_id = data.aws_vpc_endpoint_service.example.service_id
  principal_arn           = data.aws_caller_identity.current.arn
}


data "aws_vpc_endpoint_service" "example" {
  service      = "s3"
  service_type = "Gateway"
}

output "service" {
  value = data.aws_vpc_endpoint_service.example.service_name
}
