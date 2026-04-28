resource "azurerm_sentinel_data_connector_aws_s3" "pike_gen" {
  name                       = "example"
  log_analytics_workspace_id = "pike"
  aws_role_arn               = "arn:aws:iam::000000000000:role/role1"
  destination_table          = "AWSGuardDuty"
  sqs_urls                   = ["https://sqs.us-east-1.amazonaws.com/000000000000/example"]
}
