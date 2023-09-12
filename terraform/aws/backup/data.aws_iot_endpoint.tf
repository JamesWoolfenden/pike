data "aws_iot_endpoint" "pike" {}
output "endpoint" {
  value = data.aws_iot_endpoint.pike
}
