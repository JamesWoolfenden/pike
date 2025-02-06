resource "aws_pinpointsmsvoicev2_configuration_set" "pike" {
  name                 = "example-configuration-set"
  default_sender_id    = "example"
  default_message_type = "TRANSACTIONAL"

  tags = {
    pike = "permissions"
  }
}
