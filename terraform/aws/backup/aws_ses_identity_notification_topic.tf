resource "aws_ses_identity_notification_topic" "example" {
  topic_arn                = "topic"
  notification_type        = "Bounce"
  identity                 = "identity"
  include_original_headers = true
}
