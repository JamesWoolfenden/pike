resource "azurerm_sentinel_data_connector_aws_cloud_trail" "pike_gen" {
  name                       = "example"
  log_analytics_workspace_id = "pike"
  aws_role_arn               = "arn:aws:iam::000000000000:role/role1"
}
