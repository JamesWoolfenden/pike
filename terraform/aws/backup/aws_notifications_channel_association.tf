resource "aws_notifications_channel_association" "pike" {
  arn                            = aws_notificationscontacts_email_contact.pike.arn
  notification_configuration_arn = aws_notifications_notification_configuration.pike.arn

}
