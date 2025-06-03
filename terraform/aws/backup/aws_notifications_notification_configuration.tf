resource "aws_notifications_notification_configuration" "pike" {
  name        = "pike3"
  description = "Example notification configuration modified"

  tags = {
    pike        = "permissions"
    Environment = "production"
    Project     = "example"
  }
}
