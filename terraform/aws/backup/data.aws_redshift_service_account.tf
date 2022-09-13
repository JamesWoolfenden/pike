data "aws_redshift_service_account" "pike" {}

output "reds_service_account" {
  value = data.aws_redshift_service_account.pike
}
