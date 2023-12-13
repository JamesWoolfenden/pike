data "aws_iot_registration_code" "pike" {}

output "code" {
  value = data.aws_iot_registration_code.pike
}
