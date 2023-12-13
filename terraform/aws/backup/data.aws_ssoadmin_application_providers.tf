data "aws_ssoadmin_application_providers" "pike" {}

output "providers" {
  value = data.aws_ssoadmin_application_providers.pike
}
