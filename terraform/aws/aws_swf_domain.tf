resource "aws_swf_domain" "pike" {
  name                                        = "foo"
  description                                 = "Terraform SWF Domain"
  workflow_execution_retention_period_in_days = "7"
  tags = {
    pike = "permissions"
  }
}
