data "aws_billing_service_account" "pike" {}

output "acc" {
  value = data.aws_billing_service_account.pike
}
