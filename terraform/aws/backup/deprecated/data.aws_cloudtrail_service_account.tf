data "aws_cloudtrail_service_account" "example" {}

output "service_account" {
  value = data.aws_cloudtrail_service_account.example
}
