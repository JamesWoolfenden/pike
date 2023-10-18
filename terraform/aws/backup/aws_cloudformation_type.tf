resource "aws_cloudformation_type" "pike" {
  schema_handler_package = "s3://pike-680235478471/test.ps1"
  type                   = "RESOURCE"
  type_name              = "ExampleCompany::ExampleService::ExampleResource"
  logging_config {
    log_group_name = "/aws/lambda/reads3"
    log_role_arn   = "arn:aws:iam::680235478471:role/apigatewaytoclouwatch"
  }
}
