resource "aws_cloudtrail_event_data_store" "pike" {
  provider         = aws.central
  name             = "pike2"
  retention_period = 7
  tags = {
    pike    = "permissions"
    another = "tag"
  }
  termination_protection_enabled = false
}
